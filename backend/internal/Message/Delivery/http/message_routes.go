package http

import (
	"github.com/gin-gonic/gin"
)

type MessageRoutes struct {
	MessageController *MessageController
}

func NewMessageRoutes(MessageController * MessageController) *MessageRoutes {
	return &MessageRoutes{MessageController: MessageController}
}

func (m *MessageRoutes) RegisterMessageRoutes(r *gin.Engine) {
	r.POST("/send-message", m.MessageController.SendMessage)
	r.GET("/chats/:id/messages", m.MessageController.GetChatMessages)
}