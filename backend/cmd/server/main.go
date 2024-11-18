package main

import (
	"backend/cmd/setup"
	"backend/internal/config"
	"backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize services, repositories, and controllers
	authController, chatController, messageController, userController, _ := setup.Initialize()

	// Setup gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, authController, chatController, messageController, userController)

	// Run server
	r.Run(":" + config.GetPort())
}