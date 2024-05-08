package models

type Resource_Model struct {
	Inventory int32  `json:"inventory"`
	Tanks     int32  `json:"tanks"`
	Repairs   int32  `json:"repairs"`
	Player    string `json:"player"`
	World     string `json:"world"`
	Fc_id     string `json:"fcid"`
}
