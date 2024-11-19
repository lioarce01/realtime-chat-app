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
	userRepo, chatRepo, messageRepo := initializeRepositories()

	// Implement ports
	userPort, chatPort, messagePort := initializePorts(userRepo, chatRepo, messageRepo)

	// Initialize services
	webSocketService := initializeWebSocketService()
	chatService := CUseCase.NewChatService(chatPort)
	userService := UUseCase.NewUserService(userPort)
	messageService := MUseCase.NewMessageService(webSocketService, messagePort, chatPort)

	// Start WebSocket broadcasting
	go webSocketService.BroadcastMessages()

	// Initialize controllers
	authController, userController, chatController, messageController := initializeControllers(userPort, userService, chatService, messageService)

	// Register routes
	registerRoutes(router, userController, authController, chatController, messageController)

	// Return initialized components
	return authController, chatController, messageController, userController, webSocketService, router
}

func initializeRepositories() (*URepository.UserRepository, *CRepository.ChatRepository, *MRepository.MessageRepository) {
	return URepository.NewUserRepository(), &CRepository.ChatRepository{}, &MRepository.MessageRepository{}
}

func initializePorts(userRepo *URepository.UserRepository, chatRepo *CRepository.ChatRepository, messageRepo *MRepository.MessageRepository) (UPorts.UserPort, CPorts.ChatPort, MPorts.MessagePort) {
	return userRepo, chatRepo, messageRepo
}

func initializeWebSocketService() *services.WebSocketManager {
	webSocketService := services.NewWebSocketManager()
	if webSocketService == nil {
		log.Fatal("WebSocketService is nil")
	}
	return webSocketService
}

func initializeControllers(userPort UPorts.UserPort, userService *UUseCase.UserService, chatService *CUseCase.ChatService, messageService *MUseCase.MessageService) (*UserHTTP.AuthController, *UserHTTP.UserController, *ChatHTTP.ChatController, *MessageHTTP.MessageController) {
	authController := UserHTTP.NewAuthController(userPort)
	userController := UserHTTP.NewUserController(userService)
	chatController := ChatHTTP.NewChatController(chatService)
	messageController := MessageHTTP.NewMessageController(messageService)

	return authController, userController, chatController, messageController
}

func registerRoutes(router *gin.Engine, userController *UserHTTP.UserController, authController *UserHTTP.AuthController, chatController *ChatHTTP.ChatController, messageController *MessageHTTP.MessageController) {
	// Register user-related routes
	userRoutes := UserHTTP.NewUserRoutes(userController, authController)
	userRoutes.RegisterUserRoutes(router)

	// Register chat-related routes
	chatRoutes := ChatHTTP.NewChatRoutes(chatController)
	chatRoutes.RegisterChatRoutes(router)
	
	// Register message-related routes
	messageRoutes := MessageHTTP.NewMessageRoutes(messageController)
	messageRoutes.RegisterMessageRoutes(router)
}
