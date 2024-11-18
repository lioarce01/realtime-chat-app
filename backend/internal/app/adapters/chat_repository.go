package adapters

import (
	"backend/internal/app/ports"
	"backend/internal/config"
	"backend/internal/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ ports.ChatPort = &ChatRepository{}

type ChatRepository struct {}

func (r *ChatRepository) CreateChat(user1ID, user2ID primitive.ObjectID) (*models.Chat, error) {
	collection := config.DB.Collection("chats")

	var existingChat models.Chat
	err := collection.FindOne(context.TODO(), bson.M{
		"$or": []bson.M{
			{"user1_id": user1ID, "user2_id": user2ID},
			{"user1_id": user2ID, "user2_id": user1ID},
		},
	}).Decode(&existingChat)

	if err == nil {
		return nil, fmt.Errorf("chat already exists between these users")
	}

	chat := &models.Chat{
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

func (r *ChatRepository) GetChatsByUserID(userID primitive.ObjectID) ([]models.Chat, error) {
	collection := config.DB.Collection("chats")
	var chats []models.Chat

	cursor, err := collection.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"user1_id": userID},
			{"user2_id": userID},
		},
	})

	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &chats); err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *ChatRepository) FindChat(user1ID, user2ID primitive.ObjectID) (*models.Chat, error) {
    if r == nil {
        log.Println("ChatRepository is nil")
        return nil, fmt.Errorf("repository is nil")
    }

    collection := config.DB.Collection("chats")
    var chat models.Chat

    filter := bson.M{
        "$or": []bson.M{
            {"user1_id": user1ID, "user2_id": user2ID},
            {"user1_id": user2ID, "user2_id": user1ID},
        },
    }

    err := collection.FindOne(context.TODO(), filter).Decode(&chat)
    if err != nil {
        log.Printf("Error finding chat between %s and %s: %v", user1ID.Hex(), user2ID.Hex(), err)
        return nil, fmt.Errorf("chat not found")
    }

    return &chat, nil
}
