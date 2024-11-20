package usecase

import (
	"backend/config"
	CPorts "backend/internal/Domain/Chat/Ports"
	domain "backend/internal/Domain/Message/Domain"
	MPorts "backend/internal/Domain/Message/Ports"
	UPorts "backend/internal/Domain/User/Ports"
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
	UserRepo UPorts.UserPort

}

func NewMessageService(userRepo UPorts.UserPort,wsManager *services.WebSocketManager, messageRepo MPorts.MessagePort, chatService CPorts.ChatPort) *MessageService {
	return &MessageService{
		WebSocketManager: wsManager,
		MessageRepo: messageRepo,
		ChatService: chatService,
		UserRepo: userRepo,
	}
}

func (service *MessageService) SendMessage(senderID, receiverID primitive.ObjectID, content string) (*domain.Message, error) {
	chat, err := service.ChatService.FindOrCreateChat(senderID, receiverID)
	if err != nil {
		log.Println("Error finding or creating chat:", err)
		return nil, err
	}

	senderIDStr := senderID.Hex() 
	receiverIDStr := receiverID.Hex()

	sender, err := service.UserRepo.GetUserBySubOrID(senderIDStr)
	if err != nil {
		log.Println("Error retrieving sender details:", err)
		return nil, err
	}

	receiver, err := service.UserRepo.GetUserBySubOrID(receiverIDStr)
	if err != nil {
		log.Println("Error retrieving receiver details:", err)
		return nil, err
	}

	senderDetail := domain.UserDetail{
		ID:         sender.ID,
		Username:   sender.Username,
		Profile_Pic: sender.Profile_Pic,
	}

	receiverDetail := domain.UserDetail{
		ID:         receiver.ID,
		Username:   receiver.Username,
		Profile_Pic: receiver.Profile_Pic,
	}

	message := &domain.Message{
		Sender:   senderDetail,  
		Receiver: receiverDetail, 
		Content:  content,
		CreatedAt: time.Now(),
		ChatID:   chat.ID,
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
	// Retrieve enriched messages with sender and receiver details from the repository
	messages, err := service.MessageRepo.GetMessagesByChatID(chatID)
	if err != nil {
		log.Println("Error retrieving messages:", err)
		return nil, err
	}

	return messages, nil
}