package message

import (
	domain "backend/internal/Message/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessagePort interface {
	SendMessage(message *domain.Message) error
	GetMessagesByChatID(chatID primitive.ObjectID) ([]domain.Message, error)
}
