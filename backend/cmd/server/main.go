package main

import (
	"backend/cmd/setup"
	"backend/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize all dependencies and get the router
	_, _, _, _, webSocketManager, router := setup.Initialize()

	router.GET("/ws/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")
		webSocketManager.HandleWebSocket(c, userID)
	})

	// Start the server
	port := config.GetPort()
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}