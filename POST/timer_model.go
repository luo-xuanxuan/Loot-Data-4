package POST

import (
	db "LootData4/database"
	"LootData4/models"
	ws "LootData4/websocket"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Timer_Model_Handler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received body: %s\n", body)

	// Decode the JSON body into the struct
	var timer_data models.Timer_Model
	if err := json.Unmarshal(body, &timer_data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	ws.Send_Websocket_Message(body)

	db.Insert_Timer_Model(timer_data)

	// Respond with success
	w.WriteHeader(http.StatusOK)
}
