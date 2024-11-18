package ports

import (
	"backend/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatPort interface {
	CreateChat(user1ID, user2ID primitive.ObjectID) (*models.Chat, error)
	GetChatsByUserID(userID primitive.ObjectID) ([]models.Chat, error)
	FindOrCreateChat(user1ID, user2ID primitive.ObjectID) (*models.Chat, error)
}