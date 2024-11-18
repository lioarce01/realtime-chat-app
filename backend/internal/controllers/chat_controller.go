package controllers

import (
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatController struct {
	MessageService *services.MessageService
	ChatService    *services.ChatService 
}

func NewChatController(messageService *services.MessageService, chatService *services.ChatService) *ChatController {
	return &ChatController{
		MessageService: messageService,
		ChatService:    chatService,  
	}
}

func (controller *ChatController) SendMessage(c *gin.Context) {
	var messageRequest struct {
		SenderID   string `json:"sender_id"`
		ReceiverID string `json:"receiver_id"`
		Content    string `json:"content"`
	}

	if err := c.ShouldBindJSON(&messageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	senderID, err := primitive.ObjectIDFromHex(messageRequest.SenderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sender_id"})
		return
	}

	receiverID, err := primitive.ObjectIDFromHex(messageRequest.ReceiverID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver_id"})
		return
	}

	message, err := controller.MessageService.SendMessage(senderID, receiverID, messageRequest.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (controller *ChatController) CreateChat(c *gin.Context) {
	var chatRequest struct {
		User1ID string `json:"user1_id"`
		User2ID string `json:"user2_id"`
	}

	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user1ID, err := primitive.ObjectIDFromHex(chatRequest.User1ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user1_id"})
		return
	}

	user2ID, err := primitive.ObjectIDFromHex(chatRequest.User2ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user2_id"})
		return
	}

	chat, err := controller.ChatService.CreateChat(user1ID, user2ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create chat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"chat": chat})
}
