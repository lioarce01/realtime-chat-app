package services

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
	mutex     sync.Mutex
	upgrader  websocket.Upgrader
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Be cautious with this in production
			},
		},
	}
}

func (manager *WebSocketManager) HandleWebSocket(c *gin.Context) {
	conn, err := manager.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	go manager.HandleConnections(conn)
}

func (manager *WebSocketManager) HandleConnections(conn *websocket.Conn) {
	defer conn.Close()

	manager.mutex.Lock()
	manager.clients[conn] = true
	manager.mutex.Unlock()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			manager.mutex.Lock()
			delete(manager.clients, conn)
			manager.mutex.Unlock()
			break
		}

		manager.broadcast <- message
	}
}

func (manager *WebSocketManager) BroadcastMessages() {
	for {
		message := <-manager.broadcast

		manager.mutex.Lock()
		for client := range manager.clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error sending message: %v", err)
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