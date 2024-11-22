package repository

import (
	"backend/config"
	domain "backend/internal/Domain/Chat/Domain"
	ports "backend/internal/Domain/Chat/Ports"
	messageDomain "backend/internal/Domain/Message/Domain"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ ports.ChatPort = &ChatRepository{}

type ChatRepository struct{}

func NewChatRepository() *ChatRepository {
	return &ChatRepository{}
}

func (r *ChatRepository) CreateChat(user1ID, user2ID primitive.ObjectID) (*domain.Chat, error) {
	collection := config.DB.Collection("chats")

	var existingChat domain.Chat
	err := collection.FindOne(context.TODO(), bson.M{
		"$or": []bson.M{
			{"user1_id": user1ID, "user2_id": user2ID},
			{"user1_id": user2ID, "user2_id": user1ID},
		},
	}).Decode(&existingChat)

	if err == nil {
		return nil, fmt.Errorf("chat already exists between these users")
	}

	chat := &domain.Chat{
		ID:        primitive.NewObjectID(),
		User1ID:   user1ID,
		User2ID:   user2ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = collection.InsertOne(context.TODO(), chat)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (r *ChatRepository) GetChatByID(chatID primitive.ObjectID) (*domain.Chat, error) {
	chatCollection := config.DB.Collection("chats")
	messageCollection := config.DB.Collection("messages")

	var chat domain.Chat

	filter := bson.M{"_id": chatID}
	err := chatCollection.FindOne(context.TODO(), filter).Decode(&chat)
	if err != nil {
		return nil, fmt.Errorf("chat not found")
	}

	messageFilter := bson.M{
		"chat_id": chatID,
	}
	cursor, err := messageCollection.Find(context.TODO(), messageFilter)
	if err != nil {
		return nil, fmt.Errorf("error fetching messages: %v", err)
	}
	defer cursor.Close(context.TODO())

	var messages []messageDomain.Message
	for cursor.Next(context.TODO()) {
		var message messageDomain.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, fmt.Errorf("error decoding message: %v", err)
		}
		messages = append(messages, message)
	}
	chat.Messages = messages

	return &chat, nil
}

func (r *ChatRepository) GetChatsByUserID(userID primitive.ObjectID) ([]domain.Chat, error) {
	chatCollection := config.DB.Collection("chats")
	userCollection := config.DB.Collection("users")
	messageCollection := config.DB.Collection("messages")

	var chats []domain.Chat

	// Encuentra los chats del usuario
	cursor, err := chatCollection.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"user1_id": userID},
			{"user2_id": userID},
		},
	})
	if err != nil {
		log.Println("Error fetching chats:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Decodifica los chats encontrados
	if err := cursor.All(context.TODO(), &chats); err != nil {
		log.Println("Error decoding chats:", err)
		return nil, err
	}

	// Obtén los detalles de los usuarios y el último mensaje para cada chat
	for i := range chats {
		// Busca User1
		var user1 domain.UserDetail
		err := userCollection.FindOne(context.TODO(), bson.M{"_id": chats[i].User1ID}).Decode(&user1)
		if err != nil {
			log.Println("Error fetching user1 details:", err)
			continue
		}
		chats[i].User1 = &user1

		// Busca User2
		var user2 domain.UserDetail
		err = userCollection.FindOne(context.TODO(), bson.M{"_id": chats[i].User2ID}).Decode(&user2)
		if err != nil {
			log.Println("Error fetching user2 details:", err)
			continue
		}
		chats[i].User2 = &user2

		// Busca el último mensaje del chat
		var lastMessage *messageDomain.Message
		messageCursor, err := messageCollection.Find(context.TODO(), bson.M{"chat_id": chats[i].ID}, &options.FindOptions{
			Sort: bson.D{{Key: "created_at", Value: -1}},
			Limit: pointerToInt64(1),
		})
		if err != nil {
			log.Println("Error fetching last message:", err)
			continue
		}
		defer messageCursor.Close(context.TODO())

		if messageCursor.Next(context.TODO()) {
			if err := messageCursor.Decode(&lastMessage); err != nil {
				log.Println("Error decoding last message:", err)
				continue
			}
		}
		chats[i].LastMessage = lastMessage
	}

	return chats, nil
}


func (r *ChatRepository) FindOrCreateChat(user1ID, user2ID primitive.ObjectID) (*domain.Chat, error) {
	collection := config.DB.Collection("chats")
	var chat domain.Chat

	filter := bson.M{
		"$or": []bson.M{
			{"user1_id": user1ID, "user2_id": user2ID},
			{"user1_id": user2ID, "user2_id": user1ID},
		},
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&chat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			chat := domain.Chat{
				User1ID:   user1ID,
				User2ID:   user2ID,
				CreatedAt: time.Now(),
			}

			result, err := collection.InsertOne(context.TODO(), chat)
			if err != nil {
				log.Println("Error creating chat:", err)
				return nil, fmt.Errorf("error creating chat")
			}

			chat.ID = result.InsertedID.(primitive.ObjectID)
			return &chat, nil
		}

		log.Println("Error finding chat:", err)
		return nil, fmt.Errorf("chat not found")
	}

	return &chat, nil
}

func pointerToInt64(i int64) *int64 {
	return &i
}
