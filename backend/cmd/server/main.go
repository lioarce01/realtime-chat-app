package main

import (
	"backend/internal/app/adapters"
	"backend/internal/app/ports"
	"backend/internal/config"
	"backend/internal/controllers"
	"backend/internal/middlewares"
	"backend/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//Load env and connect to DB
	config.LoadEnv()
	config.ConnectDB()

	//Initialize repositories
	userRepo := &adapters.UserRepository{}
	chatRepo := &adapters.ChatRepository{}
	messageRepo := &adapters.MessageRepository{}
	
	//Implement ports to repositories
	var _ ports.UserPort = userRepo
	var _ ports.ChatPort = chatRepo
	var _ ports.MessagePort = messageRepo
	
	//Initialize services
	webSocketService := services.NewWebSocketManager()
	if webSocketService == nil {
		log.Fatal("WebSocketService is nil")
	}
	chatService := services.NewChatService(chatRepo)
	messageService := services.NewMessageService(webSocketService, chatService)
	
	//Initialize goroutine
	go webSocketService.BroadcastMessages()
	
	//Initialize controllers
	authController := &controllers.AuthController{UserPort: &adapters.UserRepository{}}
	chatController := controllers.NewChatController(messageService, chatService)

	//Setup gin router
	r := gin.Default()

	//Public routes
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	//Protected routes
	r.Use(middlewares.AuthMiddleware())
	r.POST("/send-message", chatController.SendMessage)
	r.POST("/create-chat", chatController.CreateChat) 

	//Run server
	r.Run(":" + config.GetPort())
}
