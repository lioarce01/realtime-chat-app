package adapters

import (
	"backend/internal/config"
	"backend/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type ChatService struct {}

func (s *ChatService) SendMessage(msg *models.Message) error {
	collection := config.DB.Collection("messages")
	_, err := collection.InsertOne(context.TODO(), msg)
	return err
}

func (s *ChatService) ReceiveMessages() ([]models.Message, error) {
	var messages []models.Message
	collection := config.DB.Collection("messages")
	
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var message models.Message

		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}