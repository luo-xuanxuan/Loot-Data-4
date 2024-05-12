package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Store active WebSocket clients and messages
var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

// Send a WebSocket message to all connected clients
func Send_Websocket_Message(message []byte) {
	// Broadcast to all connected clients
	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Error writing to client: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func Websocket_Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	// Register the new client
	clients[conn] = true
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()

	// Read incoming messages from the client and echo them back
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		fmt.Printf("Received: %s\n", message)
		Send_Websocket_Message(message)
	}
}
