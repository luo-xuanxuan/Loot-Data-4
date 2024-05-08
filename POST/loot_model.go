package POST

import (
	db "LootData4/database"
	"LootData4/models"
	"encoding/json"
	"net/http"
)

func Loot_Model_Handler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body into the struct
	var loot_data []models.Loot_Model
	err := json.NewDecoder(r.Body).Decode(&loot_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, loot := range loot_data {
		db.Insert_Loot_Model(loot)
		db.Update_FC_Name(loot.Fc_id, loot.World)
	}

}
