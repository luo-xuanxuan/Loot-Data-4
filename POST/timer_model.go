package POST

import (
	db "LootData4/database"
	"LootData4/models"
	ws "LootData4/websocket"
	"bytes"
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

	// Read the request body into a buffer to duplicate the stream
	body_bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read request body: %v", err), http.StatusBadRequest)
		return
	}
	// Restore the request body to its original state
	//r.Body = io.NopCloser(bytes.NewBuffer(body_bytes))

	// Log the original request body if needed
	ws.Send_Websocket_Message(body_bytes)

	// Decode the JSON body into the struct
	var timer_data []models.Timer_Model
	if err := json.NewDecoder(bytes.NewReader(body_bytes)).Decode(&timer_data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Insert each timer model into the database
	for _, timer := range timer_data {
		db.Insert_Timer_Model(timer)
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
}
