package services

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	clients   map[*websocket.Conn]string // map to associate clients with their user identifiers
	broadcast chan []byte
	mutex     sync.Mutex
	upgrader  websocket.Upgrader
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:   make(map[*websocket.Conn]string),
		broadcast: make(chan []byte),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // allow all origins (you may want to restrict this in production)
			},
		},
	}
}

// HandleWebSocket upgrades the HTTP connection to a WebSocket connection.
func (manager *WebSocketManager) HandleWebSocket(c *gin.Context) {
	conn, err := manager.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	// Set the user identifier from the URL or header if necessary
	// Example: "user_id" could be passed in as a header or URL param.
	userID := c.Param("user_id") // Assume user_id is passed as a URL param

	// Add the connection to the clients map
	manager.mutex.Lock()
	manager.clients[conn] = userID
	manager.mutex.Unlock()

	log.Printf("New WebSocket connection established for user: %s", userID)

	// Start the connection handler for the incoming messages
	manager.HandleConnections(conn)
}

// HandleConnections manages reading messages from a single WebSocket connection.
func (manager *WebSocketManager) HandleConnections(conn *websocket.Conn) {
	defer func() {
		// Cleanup on connection close
		manager.mutex.Lock()
		delete(manager.clients, conn)
		manager.mutex.Unlock()
		conn.Close()
		log.Printf("Connection closed for user: %s", manager.clients[conn])
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}

		// Forward the message to the broadcast channel
		manager.broadcast <- message
	}
}

// BroadcastMessages listens for messages on the broadcast channel and sends them to all connected clients.
func (manager *WebSocketManager) BroadcastMessages() {
	for message := range manager.broadcast {
		manager.mutex.Lock()
		for client, userID := range manager.clients {
			// Send the message to each client (could implement specific conditions for each user if needed)
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error sending message to user %s: %v", userID, err)
				client.Close()
				delete(manager.clients, client)
			}
		}
		manager.mutex.Unlock()
	}
}

// BroadcastMessage allows sending a message to all connected clients immediately.
func (manager *WebSocketManager) BroadcastMessage(message []byte) {
	manager.broadcast <- message
}
