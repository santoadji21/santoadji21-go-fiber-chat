package handler

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/santoadji21/santoadji21-go-fiber-chat/internal/db"
	"github.com/santoadji21/santoadji21-go-fiber-chat/internal/redis"
	"github.com/santoadji21/santoadji21-go-fiber-chat/pkg/models"
)

// Global variables for managing WebSocket connections
var (
	connections = make(map[*websocket.Conn]bool)
	connMutex   = sync.Mutex{}
)

// WebSocketChatHandler handles incoming WebSocket connections
func WebSocketChatHandler(c *websocket.Conn) {
	connMutex.Lock()
	connections[c] = true
	connMutex.Unlock()

	defer func() {
		connMutex.Lock()
		delete(connections, c)
		connMutex.Unlock()
		c.Close()
	}()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		// Unmarshal the message into a Message struct
		var message models.Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		// Save message to the database
		if err := db.GetDB().Create(&message).Error; err != nil {
			log.Println("Failed to save message:", err)
		}

		// Publish message to Redis
		ctx := context.Background()
		if err := redis.GetRedis().Publish(ctx, "chatroom:"+message.Room, msg).Err(); err != nil {
			log.Println("Failed to publish message to Redis:", err)
		}

		// Broadcast message to all connected clients
		broadcastMessage(msg)
	}
}

// broadcastMessage sends the message to all connected WebSocket clients
func broadcastMessage(msg []byte) {
	connMutex.Lock()
	defer connMutex.Unlock()

	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Error broadcasting message:", err)
			delete(connections, conn)
		}
	}
}
