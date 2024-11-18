package services

import (
	"backend/internal/app/ports"
	"backend/internal/config"
	"backend/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type MessageService struct {
    WebSocketManager *WebSocketManager
    MessageRepo ports.MessagePort
    ChatService ports.ChatPort
}

func NewMessageService(ws *WebSocketManager, messageRepo ports.MessagePort, chatService ports.ChatPort) *MessageService {
    return &MessageService{
        WebSocketManager: ws,
        MessageRepo:      messageRepo,
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

func (service *MessageService) GetMessagesByChatID(chatID primitive.ObjectID) ([]models.Message, error) {
    log.Printf("Fetching messages for chatID: %s", chatID.Hex()) 

    messages, err := service.MessageRepo.GetMessagesByChatID(chatID)
    if err != nil {
        log.Printf("Error fetching messages from repository: %v", err)
        return nil, err
    }

    log.Printf("Found %d messages for chatID: %s", len(messages), chatID.Hex()) 
    return messages, nil
}
