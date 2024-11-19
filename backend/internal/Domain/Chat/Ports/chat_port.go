package ports

import (
	domain "backend/internal/Domain/Chat/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatPort interface {
	CreateChat(user1ID, user2ID primitive.ObjectID) (*domain.Chat, error)
	GetChatsByUserID(userID primitive.ObjectID) ([]domain.Chat, error)
	FindOrCreateChat(user1ID, user2ID primitive.ObjectID) (*domain.Chat, error)
}
