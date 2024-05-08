package models

type Timer_Model struct {
	Return_time int64  `json:"return_time"`
	Fc_id       string `json:"fcid"`
	Name        string `json:"name"`
	Sub_id      int32  `json:"sub_id"`
}
