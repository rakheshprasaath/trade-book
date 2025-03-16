package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// WebhookPayload represents the structure of the incoming webhook payload
type WebhookPayload map[string]interface{}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload WebhookPayload

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Log the received payload
	fmt.Println("Webhook received:")
	jsonPayload, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println(string(jsonPayload))

	// Simulate a delay before responding
	time.Sleep(1 * time.Second) // 1 second delay

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Webhook received successfully after delay"})
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	port := ":300" // Default port
	fmt.Printf("Webhook server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
