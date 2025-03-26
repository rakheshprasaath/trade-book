package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader to handle WebSocket connections (upgrades HTTP to WebSocket)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (for testing purposes)
	},
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP request to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")

	// Continuously read and print messages from the WebSocket client
	for {
		var msg map[string]interface{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading JSON from WebSocket:", err)
			break
		}

		// Print the received message
		fmt.Println("Received message:", msg)

		// Echo the received message back to the client (optional)
		err = conn.WriteJSON(map[string]string{
			"message": "Message received",
		})
		if err != nil {
			log.Println("Error writing JSON to WebSocket:", err)
			break
		}
	}
}

// func main() {
// 	http.HandleFunc("/ws", WebSocketHandler)

// 	port := ":8080"
// 	fmt.Println("WebSocket server started on ws://localhost" + port)
// 	log.Fatal(http.ListenAndServe(port, nil))
// }
