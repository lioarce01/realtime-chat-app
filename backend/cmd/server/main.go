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
	userService := services.NewUserService(userRepo)
	
	//Initialize goroutine
	go webSocketService.BroadcastMessages()
	
	//Initialize controllers
	authController := &controllers.AuthController{UserPort: &adapters.UserRepository{}}
	chatController := controllers.NewChatController(messageService, chatService)
	userController := controllers.NewUserController(userService)

	//Setup gin router
	r := gin.Default()

	//Public User routes
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserByID)

	//Public authentication routes
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	//Protected chat routes
	r.Use(middlewares.AuthMiddleware())
	r.POST("/send-message", chatController.SendMessage)
	r.POST("/create-chat", chatController.CreateChat) 

	//Run server
	r.Run(":" + config.GetPort())
}
