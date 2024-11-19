package repository

import (
	domain "backend/internal/Chat/Domain"
	ports "backend/internal/Chat/Ports"
	"backend/internal/config"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *ChatRepository) GetChatsByUserID(userID primitive.ObjectID) ([]domain.Chat, error) {
	collection := config.DB.Collection("chats")
	var chats []domain.Chat

	cursor, err := collection.Find(context.TODO(), bson.M{
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

	if err := cursor.All(context.TODO(), &chats); err != nil {
		log.Println("Error decoding chats:", err)
		return nil, err
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
		log.Println("Error finding chat:", err)
		return nil, fmt.Errorf("chat not found")
	}

	return &chat, nil
}
