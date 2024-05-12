package POST

import (
	db "LootData4/database"
	"LootData4/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func in_list(i int, arr []int) bool {
	for _, v := range arr {
		if v == i {
			return true
		}
	}
	return false
}

func Loot_Model_Handler(w http.ResponseWriter, r *http.Request) {
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
	var loot_data []models.Loot_Model
	if err := json.Unmarshal(body, &loot_data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	for _, loot := range loot_data {
		if in_list(loot.Item_id, []int{22500, 22501, 22502, 22503, 22504, 22505, 22506, 22507}) {
			db.Insert_Loot_Model(loot)
			db.Update_FC_Name(loot.Fc_id, loot.World)
		}
	}

}
