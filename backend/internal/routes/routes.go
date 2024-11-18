package routes

import (
	"backend/internal/controllers"
	"backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authController *controllers.AuthController, 
	chatController *controllers.ChatController, 
	messageController *controllers.MessageController, 
	userController *controllers.UserController) {

	// Public User routes
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.GET("/users/:id/chats", chatController.GetUserChats)

	// Public authentication routes
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	// Protected routes
	r.Use(middlewares.AuthMiddleware())

	// Message routes
	r.POST("/send-message", messageController.SendMessage)
	r.GET("/chats/:id/messages", messageController.GetChatMessages)

	// Chat routes
	r.POST("/create-chat", chatController.CreateChat)
}
