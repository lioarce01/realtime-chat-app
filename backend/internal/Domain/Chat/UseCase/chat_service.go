package services

import (
	domain "backend/internal/Domain/Chat/Domain"
	ports "backend/internal/Domain/Chat/Ports"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatService struct {
	ChatRepo ports.ChatPort
}

func NewChatService(chatRepo ports.ChatPort) *ChatService {
	if chatRepo == nil {
		log.Fatal("ChatRepository is nil")
	}

	return &ChatService{
		ChatRepo: chatRepo,
	}
}

func (s *ChatService) CreateChat(user1ID, user2ID primitive.ObjectID) (*domain.Chat, error) {
	return s.ChatRepo.CreateChat(user1ID, user2ID)
}

func (s *ChatService) GetUserChats(userID primitive.ObjectID) ([]domain.Chat, error) {
	return s.ChatRepo.GetChatsByUserID(userID)
}

func (s *ChatService) GetChatByID(chatID primitive.ObjectID) (*domain.Chat, error) {
	return s.ChatRepo.GetChatByID(chatID)
}

func (s *ChatService) FindOrCreateChat(user1ID, user2ID primitive.ObjectID) (*domain.Chat, error) {
	chat, err := s.ChatRepo.FindOrCreateChat(user1ID, user2ID)
	if err == nil {
		return chat, nil
	}

	if err.Error() == "chat not found" {
		return s.ChatRepo.CreateChat(user1ID, user2ID)
	}

	return nil, fmt.Errorf("unexpected error: %v", err)
}

func (s *ChatService) DeleteChatByID(chatID primitive.ObjectID) error {
	chat, err := s.ChatRepo.GetChatByID(chatID)
	if err != nil {
		return err
	}
	if chat == nil {
		return errors.New("chat not found")
	}

	return s.ChatRepo.DeleteChatByID(chatID)
}
