package POST

import (
	db "LootData4/database"
	"LootData4/models"
	"encoding/json"
	"net/http"
)

func Resource_Model_Handler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body into the struct
	var resource_data []models.Resource_Model
	err := json.NewDecoder(r.Body).Decode(&resource_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, resource := range resource_data {
		db.Insert_Resource_Model(resource)
		db.Update_FC_Name(resource.Fc_id, resource.World)
	}
}
