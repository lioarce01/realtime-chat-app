package http

import (
	"github.com/gin-gonic/gin"
)

type ChatRoutes struct {
	ChatController *ChatController
}

func NewChatRoutes(chatController *ChatController) * ChatRoutes {
	return &ChatRoutes{ChatController: chatController}
}

func (c *ChatRoutes) RegisterChatRoutes(r *gin.Engine) {
	r.POST("/create-chat", c.ChatController.CreateChat)
	r.GET("/users/:id/chats",c.ChatController.GetUserChats)
}