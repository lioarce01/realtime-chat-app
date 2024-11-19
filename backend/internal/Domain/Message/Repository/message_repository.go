package repository

import (
	"backend/config"
	domain "backend/internal/Domain/Message/Domain"
	ports "backend/internal/Domain/Message/Ports"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepository struct{}

var _ ports.MessagePort = &MessageRepository{}

func (r *MessageRepository) SendMessage(message *domain.Message) error {
	collection := config.DB.Collection("messages")
	message.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), message)
	return err
}

func (r *MessageRepository) GetMessagesByChatID(chatID primitive.ObjectID) ([]domain.Message, error) {
	collection := config.DB.Collection("messages")
	var messages []domain.Message

	cursor, err := collection.Find(context.TODO(), bson.M{"chat_id": chatID})
	if err != nil {
		log.Printf("Error querying messages: %v", err)
		return nil, err
	}

	if err := cursor.All(context.TODO(), &messages); err != nil {
		log.Printf("Error retrieving messages from cursor: %v", err)
		return nil, err
	}

	return messages, nil
}
