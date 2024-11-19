package setup

import (
	"backend/config"
	ChatHTTP "backend/internal/Domain/Chat/Delivery/http"
	CPorts "backend/internal/Domain/Chat/Ports"
	CRepository "backend/internal/Domain/Chat/Repository"
	CUseCase "backend/internal/Domain/Chat/UseCase"
	MessageHTTP "backend/internal/Domain/Message/Delivery/http"
	MPorts "backend/internal/Domain/Message/Ports"
	MRepository "backend/internal/Domain/Message/Repository"
	MUseCase "backend/internal/Domain/Message/UseCase"
	UserHTTP "backend/internal/Domain/User/Delivery/http"
	UPorts "backend/internal/Domain/User/Ports"
	URepository "backend/internal/Domain/User/Repository"
	UUseCase "backend/internal/Domain/User/UseCase"
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