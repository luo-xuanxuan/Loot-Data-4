package models

type World_Status struct {
	Name              string                 `json:"name"`
	Free_Company_List []*Free_Company_Status `json:"free_company_list"`
}

type Free_Company_Status struct {
	ID               string                `json:"id"`
	Name             string                `json:"name"`
	Tanks            int                   `json:"tanks"`
	Repairs          int                   `json:"repairs"`
	Submersible_List []*Submersible_Status `json:"submersible_list"`
}

type Submersible_Status struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Return_Time int64  `json:"return_time"`
}
