package controllers

import (
	"backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {return true},
}

type ChatController struct {
	WebSocketService *services.WebSocketManager
}

func NewChatController(webSocketService *services.WebSocketManager) *ChatController {
	return &ChatController{WebSocketService: webSocketService}
}

func (controller *ChatController) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
		return
	}

	go controller.WebSocketService.HandleConnections(conn)
}

func (controller *ChatController) SendMessage(c *gin.Context) {
	var message struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message"})
		return
	}

	controller.WebSocketService.BroadcastMessage([]byte(message.Content))

	c.JSON(http.StatusOK, gin.H{"status": "Message sent"})
}
