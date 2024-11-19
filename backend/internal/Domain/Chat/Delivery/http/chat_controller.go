package http

import (
	usecase "backend/internal/Domain/Chat/UseCase"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatController struct {
	ChatService *usecase.ChatService
}

func NewChatController(chatService *usecase.ChatService) *ChatController {
	return &ChatController{
		ChatService: chatService,
	}
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

func (controller *ChatController) GetUserChats(c *gin.Context) {
	userIDParam := c.Param("id")

	if len(userIDParam) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	chats, err := controller.ChatService.GetUserChats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve chats"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"chats": chats})
}
