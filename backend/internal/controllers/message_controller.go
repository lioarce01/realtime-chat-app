package controllers

import (
	"backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageController struct {
	MessageService *services.MessageService
}

func NewMessageController(messageService *services.MessageService) *MessageController {
	return &MessageController{
		MessageService: messageService,
	}
}

func (controller *MessageController) SendMessage(c *gin.Context) {
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

func (controller *MessageController) GetChatMessages(c *gin.Context) {
    chatIDParam := c.Param("id")
    log.Printf("Received chatID: %s", chatIDParam) 

    chatID, err := primitive.ObjectIDFromHex(chatIDParam)
    if err != nil {
        log.Printf("Error parsing chat ID: %v", err) 
        c.JSON(http.StatusBadRequest, gin.H{"error": "Chat ID must be a 24-character hex string."})
        return
    }

    messages, err := controller.MessageService.GetMessagesByChatID(chatID)
    if err != nil {
        log.Printf("Error retrieving messages: %v", err) 
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve messages"})
        return
    }

    if len(messages) == 0 {
        log.Println("No messages found for chat ID:", chatID.Hex()) 
    }

    c.JSON(http.StatusOK, gin.H{"messages": messages})
}
