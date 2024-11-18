package services

import (
	"backend/internal/config"
	"backend/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type MessageService struct {
	WebSocketManager *WebSocketManager
	ChatService *ChatService
}

func NewMessageService(ws *WebSocketManager, chatService *ChatService) *MessageService {
    return &MessageService{
        WebSocketManager: ws,
        ChatService:      chatService,  
    }
}

func (service *MessageService) SendMessage(senderID, receiverID primitive.ObjectID, content string) (*models.Message, error) {

    if service.ChatService == nil {
        log.Println("ChatService is nil")
    }

    chat, err := service.ChatService.FindOrCreateChat(senderID, receiverID)
    if err != nil {
        log.Println("Error finding or creating chat:", err)
        return nil, err
    }

    message := &models.Message{
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
