package db

import (
	"LootData4/models"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Select_World_Statuses() []*models.World_Status {

	fc_to_world := make(map[string]string, 0)

	worlds := make([]string, 0)

	world_statuses := make(map[string]models.World_Status)

	db, err := sql.Open("sqlite3", db_path+"?_busy_timeout=5000")
	if err != nil {
		log.Print("opening")
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM FC_World")
	if err != nil {
		log.Print("query fc_world")
		log.Fatal(err)
	}

	for rows.Next() {
		var fc_id string
		var world string
		var name string
		var time int64
		err = rows.Scan(&fc_id, &world, &name, &time)
		if err != nil {
			log.Print("scan fc_world")
			log.Fatal(err)
		}

		//name = Update_FC_Name(fc_id, world)

		world_status, exists := world_statuses[world]

		if !exists {
			worlds = append(worlds, world)
			world_status = models.World_Status{
				Name: world,
			}
		}

		world_status.Free_Company_List = append(world_status.Free_Company_List, &models.Free_Company_Status{
			ID:   fc_id,
			Name: name,
		})

		world_statuses[world] = world_status

		fc_to_world[fc_id] = world
	}
	rows.Close()

	rows, err = db.Query("SELECT * FROM Submersible_Resources")
	if err != nil {
		log.Print("query_resources")
		log.Fatal(err)
	}

	for rows.Next() {
		var fc_id string
		var tanks int
		var repairs int
		err = rows.Scan(&fc_id, &tanks, &repairs)
		if err != nil {
			log.Print("scan resources")
			log.Fatal(err)
		}

		world := fc_to_world[fc_id]

		for _, fc := range world_statuses[world].Free_Company_List {
			if fc.ID == fc_id {
				fc.Tanks = tanks
				fc.Repairs = repairs
				break
			}
		}

	}

	rows.Close()

	rows, err = db.Query("SELECT * FROM Submersible_Timers")
	if err != nil {
		log.Print("query timers")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var fc_id string
		var sub_name_1 string
		var return_time_1 int64
		var sub_name_2 string
		var return_time_2 int64
		var sub_name_3 string
		var return_time_3 int64
		var sub_name_4 string
		var return_time_4 int64

		err = rows.Scan(&fc_id, &sub_name_1, &return_time_1, &sub_name_2, &return_time_2, &sub_name_3, &return_time_3, &sub_name_4, &return_time_4)
		if err != nil {
			log.Print("scan timers")
			log.Fatal(err)
		}

		world := fc_to_world[fc_id]

		for _, fc := range world_statuses[world].Free_Company_List {
			if fc.ID == fc_id {
				fc.Submersible_List = []*models.Submersible_Status{
					{
						ID:          1,
						Name:        sub_name_1,
						Return_Time: return_time_1,
					},
					{
						ID:          2,
						Name:        sub_name_2,
						Return_Time: return_time_2,
					},
					{
						ID:          3,
						Name:        sub_name_3,
						Return_Time: return_time_3,
					},
					{
						ID:          4,
						Name:        sub_name_4,
						Return_Time: return_time_4,
					},
				}
				break
			}
		}

	}

	response := make([]*models.World_Status, 0)
	for _, w := range worlds {
		v := world_statuses[w]
		list := make([]*models.Free_Company_Status, 0)
		for _, f := range v.Free_Company_List {
			if len(f.Submersible_List) > 0 {
				list = append(list, f)
			}

		}
		v.Free_Company_List = list

		response = append(response, &v)

	}

	return response
}
