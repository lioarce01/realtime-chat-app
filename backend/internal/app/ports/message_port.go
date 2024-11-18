package ports

import (
	"backend/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessagePort interface {
	SendMessage(message *models.Message) error
	GetMessagesByChatID(chatID primitive.ObjectID) ([]models.Message, error)
}