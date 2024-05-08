package GET

import (
	db "LootData4/database"
	"encoding/json"
	"log"
	"net/http"
)

func World_Status_Model_Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := db.Select_World_Statuses()

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
