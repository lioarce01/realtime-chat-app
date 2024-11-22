package services

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
	ChatID     string `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
}

type WebSocketManager struct {
	clients   map[*websocket.Conn]string
	broadcast chan Message
	mutex     sync.RWMutex
	upgrader  websocket.Upgrader
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:   make(map[*websocket.Conn]string),
		broadcast: make(chan Message),
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

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		manager.broadcast <- msg
	}
}

func (manager *WebSocketManager) BroadcastMessages() {
	for message := range manager.broadcast {
		manager.mutex.RLock()
		for client, userID := range manager.clients {
			if userID == message.SenderID || userID == message.ReceiverID {
				err := client.WriteJSON(message)
				if err != nil {
					log.Printf("Error sending message to user %s: %v", userID, err)
					client.Close()
					manager.mutex.RUnlock()
					manager.mutex.Lock()
					delete(manager.clients, client)
					manager.mutex.Unlock()
					manager.mutex.RLock()
				}
			}
		}
		manager.mutex.RUnlock()
	}
}

func (manager *WebSocketManager) BroadcastMessage(message Message) {
	manager.broadcast <- message
}