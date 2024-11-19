package setup

import (
	ChatHTTP "backend/internal/Chat/Delivery/http"
	CPorts "backend/internal/Chat/Ports"
	CRepository "backend/internal/Chat/Repository"
	CUseCase "backend/internal/Chat/UseCase"
	MessageHTTP "backend/internal/Message/Delivery/http"
	MPorts "backend/internal/Message/Ports"
	MRepository "backend/internal/Message/Repository"
	MUseCase "backend/internal/Message/UseCase"
	UserHTTP "backend/internal/User/Delivery/http"
	UPorts "backend/internal/User/Ports"
	URepository "backend/internal/User/Repository"
	UUseCase "backend/internal/User/UseCase"
	"backend/internal/config"
	"backend/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize() (*UserHTTP.AuthController, *ChatHTTP.ChatController, *MessageHTTP.MessageController, *UserHTTP.UserController, *services.WebSocketManager, *gin.Engine) {
	// Load environment variables and connect to the database
	config.LoadEnv()
	config.ConnectDB()

	// Initialize Gin router
	router := gin.Default()

	// Initialize repositories
	userRepo := URepository.NewUserRepository()
	chatRepo := &CRepository.ChatRepository{}
	messageRepo := &MRepository.MessageRepository{}

	// Implement UserPort interface with UserRepository
	var userPort UPorts.UserPort = userRepo
	var chatPort CPorts.ChatPort = chatRepo
	var messagePort MPorts.MessagePort = messageRepo

	// Initialize services
	webSocketService := services.NewWebSocketManager()
	if webSocketService == nil {
		log.Fatal("WebSocketService is nil")
	}

	chatService := CUseCase.NewChatService(chatPort)
	userService := UUseCase.NewUserService(userPort)
	messageService := MUseCase.NewMessageService(webSocketService, messagePort, chatPort)

	// Start the goroutine for broadcasting messages via WebSocket
	go webSocketService.BroadcastMessages()

	// Initialize controllers
	authController := UserHTTP.NewAuthController(userPort)
	userController := UserHTTP.NewUserController(userService)
	chatController := ChatHTTP.NewChatController(chatService)
	messageController := MessageHTTP.NewMessageController(messageService)

	// Register user-related routes
	userRoutes := UserHTTP.NewUserRoutes(userController, authController)
	userRoutes.RegisterUserRoutes(router)

	chatRoutes := ChatHTTP.NewChatRoutes(chatController)
	chatRoutes.RegisterChatRoutes(router)
	
	messageRoutes := MessageHTTP.NewMessageRoutes(messageController)
	messageRoutes.RegisterMessageRoutes(router)

	// Return the initialized controllers, services, and router
	return authController, chatController, messageController, userController, webSocketService, router
}