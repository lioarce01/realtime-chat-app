package usecase

import (
	"backend/config"
	CPorts "backend/internal/Domain/Chat/Ports"
	domain "backend/internal/Domain/Message/Domain"
	MPorts "backend/internal/Domain/Message/Ports"
	"backend/internal/services"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type MessageService struct {
	WebSocketManager *services.WebSocketManager
	MessageRepo MPorts.MessagePort
	ChatService CPorts.ChatPort
}

func NewMessageService(wsManager *services.WebSocketManager, messageRepo MPorts.MessagePort, chatService CPorts.ChatPort) *MessageService {
	return &MessageService{
		WebSocketManager: wsManager,
		MessageRepo: messageRepo,
		ChatService: chatService,
	}
}

func (service *MessageService) SendMessage(senderID, receiverID primitive.ObjectID, content string) (*domain.Message, error) {

	chat, err := service.ChatService.FindOrCreateChat(senderID, receiverID)
	if err != nil {
		log.Println("Error finding or creating chat:", err)
		return nil, err
	}

	message := &domain.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		CreatedAt:  time.Now(),
		ChatID:     chat.ID,
	}

	collection := config.DB.Collection("messages")
	result, err := collection.InsertOne(context.TODO(), message)
	if err != nil {
		log.Println("Error saving message:", err)
		return nil, err
	}

	message.ID = result.InsertedID.(primitive.ObjectID)

	service.WebSocketManager.BroadcastMessage([]byte(content))

	return message, nil
}

func (service *MessageService) GetMessagesByChatID(chatID primitive.ObjectID) ([]domain.Message, error) {
	return service.MessageRepo.GetMessagesByChatID(chatID)
}
