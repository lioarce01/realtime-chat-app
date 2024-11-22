package services

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	clients   map[*websocket.Conn]string
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
				return true 
			},
		},
	}
}

func (manager *WebSocketManager) HandleWebSocket(c *gin.Context, userID string) {
	conn, err := manager.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	manager.mutex.Lock()
	manager.clients[conn] = userID
	manager.mutex.Unlock()

	log.Printf("New WebSocket connection established for user: %s", userID)

	manager.handleConnections(conn)
}

func (manager *WebSocketManager) handleConnections(conn *websocket.Conn) {
	defer func() {
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

		manager.broadcast <- message
	}
}

func (manager *WebSocketManager) BroadcastMessages() {
	for message := range manager.broadcast {
		manager.mutex.Lock()
		for client, userID := range manager.clients {
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

func (manager *WebSocketManager) BroadcastMessage(message []byte) {
	manager.broadcast <- message
}
