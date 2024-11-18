package setup

import (
	"backend/internal/app/adapters"
	"backend/internal/app/ports"
	"backend/internal/config"
	"backend/internal/controllers"
	"backend/internal/services"
	"log"
)

// Initialize all dependencies like repositories, services, and controllers
func Initialize() (*controllers.AuthController, *controllers.ChatController, *controllers.MessageController, *controllers.UserController, *services.WebSocketManager) {
	// Load environment and connect to DB
	config.LoadEnv()
	config.ConnectDB()

	// Initialize repositories
	userRepo := &adapters.UserRepository{}
	chatRepo := &adapters.ChatRepository{}
	messageRepo := &adapters.MessageRepository{}

	// Implement ports to repositories
	var _ ports.UserPort = userRepo
	var _ ports.ChatPort = chatRepo
	var _ ports.MessagePort = messageRepo

	// Initialize services
	webSocketService := services.NewWebSocketManager()
	if webSocketService == nil {
		log.Fatal("WebSocketService is nil")
	}
	chatService := services.NewChatService(chatRepo)
	messageService := services.NewMessageService(webSocketService, messageRepo, chatRepo)
	userService := services.NewUserService(userRepo)

	// Start the goroutine for broadcasting messages
	go webSocketService.BroadcastMessages()

	// Initialize controllers
	authController := &controllers.AuthController{UserPort: userRepo}
	chatController := controllers.NewChatController(chatService)
	messageController := controllers.NewMessageController(messageService)
	userController := controllers.NewUserController(userService)

	// Return the initialized controllers and WebSocketManager
	return authController, chatController, messageController, userController, webSocketService
}
