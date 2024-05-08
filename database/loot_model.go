package db

import (
	"LootData4/models"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Select_Report_Map(days int64) ([]*models.Loot_Model, error) {

	report := make([]*models.Loot_Model, 0)

	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Submersible_Loot WHERE time > ?", time.Now().Unix()-(days*86400))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var loot_row models.Loot_Model

		err = rows.Scan(&loot_row.Timestamp, &loot_row.Fc_id, &loot_row.Sub_id, &loot_row.Player, &loot_row.World, &loot_row.Item_id, &loot_row.Quantity)
		if err != nil {
			return nil, err
		}

		report = append(report, &loot_row)

	}

	return report, nil

}
