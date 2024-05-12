package POST

import (
	db "LootData4/database"
	"LootData4/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Resource_Model_Handler(w http.ResponseWriter, r *http.Request) {
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
	var resource_data models.Resource_Model
	if err := json.Unmarshal(body, &resource_data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}
	db.Insert_Resource_Model(resource_data)
	db.Update_FC_Name(resource_data.Fc_id, resource_data.World)
}
