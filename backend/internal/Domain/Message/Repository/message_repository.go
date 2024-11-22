package repository

import (
	"backend/config"
	domain "backend/internal/Domain/Message/Domain"
	ports "backend/internal/Domain/Message/Ports"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ ports.MessagePort = &MessageRepository{}

type MessageRepository struct{}

func NewMessageRepository() *MessageRepository{
	return &MessageRepository{}
}

func (r *MessageRepository) SendMessage(message *domain.Message) error {
	collection := config.DB.Collection("messages")
	message.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), message)
	return err
}

func (r *MessageRepository) GetMessagesByChatID(chatID primitive.ObjectID) ([]domain.Message, error) {
    messageCollection := config.DB.Collection("messages")

    filter := bson.M{"chat_id": chatID}
    cursor, err := messageCollection.Find(context.TODO(), filter)
    if err != nil {
        return nil, fmt.Errorf("error fetching messages: %v", err)
    }
    defer cursor.Close(context.TODO())

    var messages []domain.Message
    for cursor.Next(context.TODO()) {
        var message domain.Message
        if err := cursor.Decode(&message); err != nil {
            return nil, fmt.Errorf("error decoding message: %v", err)
        }

        messages = append(messages, message)
    }

    if err := cursor.Err(); err != nil {
        return nil, fmt.Errorf("cursor error: %v", err)
    }

    return messages, nil
}