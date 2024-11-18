package services

import (
	"backend/internal/app/adapters"
	"backend/internal/app/ports"
	"backend/internal/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatService struct {
	ChatRepo ports.ChatPort
}

func NewChatService(chatRepo *adapters.ChatRepository) *ChatService {
	if chatRepo == nil {
		log.Fatal("ChatRepository is nil")
	}

	return &ChatService{
		ChatRepo: chatRepo,
	}
}

func (s *ChatService) CreateChat(user1ID, user2ID primitive.ObjectID) (*models.Chat, error) {
	return s.ChatRepo.CreateChat(user1ID, user2ID)
}

func (s *ChatService) GetUserChats(userID primitive.ObjectID) ([]models.Chat, error) {
	return s.ChatRepo.GetChatsByUserID(userID)
}

func (s *ChatService) FindOrCreateChat(user1ID, user2ID primitive.ObjectID) (*models.Chat, error) {
    chat, err := s.ChatRepo.FindOrCreateChat(user1ID, user2ID)
    if err == nil {
        return chat, nil
    }

    if err.Error() == "chat not found" {
        return s.ChatRepo.CreateChat(user1ID, user2ID)
    }

    return nil, fmt.Errorf("unexpected error: %v", err)
}
