package main

import (
	"backend/cmd/setup"
	"backend/config"
	"log"
)

func main() {
	// Initialize all dependencies and get the router
	_, _, _, _, _, router := setup.Initialize()

	// Start the server
	port := config.GetPort()
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}
