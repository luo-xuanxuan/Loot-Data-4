package POST

import (
	db "LootData4/database"
	"LootData4/models"
	"encoding/json"
	"net/http"
)

func Timer_Model_Handler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body into the struct
	var timer_data []models.Timer_Model
	err := json.NewDecoder(r.Body).Decode(&timer_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, resource := range timer_data {
		db.Insert_Timer_Model(resource)
	}
}
