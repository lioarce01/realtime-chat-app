package main

import (
	"backend/internal/app/adapters"
	"backend/internal/config"
	"backend/internal/controllers"
	"backend/internal/middlewares"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	webSocketService := services.NewWebSocketManager()

	go webSocketService.BroadcastMessages()

	config.ConnectDB()

	userRepo := &adapters.UserRepository{}
	authController := &controllers.AuthController{UserPort: userRepo}
	chatController := controllers.NewChatController(webSocketService)

	r := gin.Default()

	// Public routes
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	// Authenticated routes
	r.Use(middlewares.AuthMiddleware())
	r.POST("/send-message", chatController.SendMessage)
	r.GET("/ws", chatController.HandleWebSocket)

	r.Run(":" + config.GetPort())
}
