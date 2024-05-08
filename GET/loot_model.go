package GET

import (
	db "LootData4/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Loot_Model_Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	queryValues := r.URL.Query()
	days := queryValues.Get("days")
	if days == "" {
		// Respond with an error if 'days' parameter is missing
		http.Error(w, "Missing 'days' query parameter", http.StatusBadRequest)
		return
	}

	daysi, err := strconv.Atoi(days)
	if err != nil {
		log.Printf("Error converting days to int: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response, err := db.Select_Report_Map(int64(daysi))
	if err != nil {
		log.Printf("Error selecting report from DB: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
