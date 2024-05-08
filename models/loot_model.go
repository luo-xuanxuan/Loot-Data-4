package models

type Loot_Model struct {
	Timestamp int64  `json:"time"`
	Fc_id     string `json:"fcid"`
	Sub_id    string `json:"sub_id"`
	Player    string `json:"player"`
	World     string `json:"world"`
	Sector_id int    `json:"sector_id"`
	Item_id   int    `json:"item_id"`
	Quantity  int    `json:"quantity"`
}
