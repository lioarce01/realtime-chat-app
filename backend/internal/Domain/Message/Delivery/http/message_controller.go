package http

import (
	usecase "backend/internal/Domain/Message/UseCase"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageController struct {
	MessageService *usecase.MessageService
}

func NewMessageController(messageService *usecase.MessageService) *MessageController {
	return &MessageController{
		MessageService: messageService,
	}
}

func (controller *MessageController) SendMessage(c *gin.Context) {
	var messageRequest struct {
		SenderID   string `json:"sender_id" binding:"required"`
		ReceiverID string `json:"receiver_id" binding:"required"`
		Content    string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&messageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (c *MessageController) GetChatMessages(g *gin.Context) {
    chatIDParam := g.Param("id")
    chatID, err := primitive.ObjectIDFromHex(chatIDParam)
    if err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
        return
    }

    messages, err := c.MessageService.GetMessagesByChatID(chatID)
    if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
        return
    }

    g.JSON(http.StatusOK, gin.H{"messages": messages})
}