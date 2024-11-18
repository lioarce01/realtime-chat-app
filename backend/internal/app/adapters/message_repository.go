package adapters

import (
	"backend/internal/app/ports"
	"backend/internal/config"
	"backend/internal/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepository struct{}

var _ ports.MessagePort = &MessageRepository{}

func (r *MessageRepository) SendMessage(message *models.Message) error {
	collection := config.DB.Collection("messages")
	message.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), message)
	return err
}

func (r *MessageRepository) GetMessagesByChatID(chatID primitive.ObjectID) ([]models.Message, error) {
	collection := config.DB.Collection("messages")
	var messages []models.Message

	cursor, err := collection.Find(context.TODO(), bson.M{"chat_id": chatID})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}