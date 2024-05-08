package db

import (
	"LootData4/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db_path = "./data.db"

func init() {

	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS FC_World (fc text PRIMARY KEY, world text, name text, updated integer)")
	statement.Exec()

	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS Submersible_Loot (time integer, fc text, sub integer, player text, world text, item integer, quantity integer)")
	statement.Exec()

	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS Submersible_Resources (fc text PRIMARY KEY, tanks integer, repairs integer)")
	statement.Exec()

	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS Submersible_Timers (fc text PRIMARY KEY, sub_1 text, return_1 integer, sub_2 text, return_2 integer, sub_3 text, return_3 integer, sub_4 text, return_4 integer)")
	statement.Exec()
}

func Insert_Loot_Model(loot_model models.Loot_Model) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO Submersible_Loot VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(loot_model.Timestamp, loot_model.Fc_id, loot_model.Sub_id, loot_model.Player, loot_model.World, loot_model.Item_id, loot_model.Quantity)
	if err != nil {
		log.Fatal(err)
	}
}

func Insert_Resource_Model(resource_model models.Resource_Model) {

	// Inventory = 0 is player
	// Inventory = 1 is FC Chest
	// We only care about FC Chest for now, so ignore player updates
	if resource_model.Inventory == 0 {
		return
	}

	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, err := db.Prepare("REPLACE INTO Submersible_Resources (fc, tanks, repairs) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(resource_model.Fc_id, resource_model.Tanks, resource_model.Repairs)
	if err != nil {
		log.Fatal(err)
	}

}

func Insert_Timer_Model(timer_model models.Timer_Model) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sub_col := fmt.Sprintf("sub_%d", timer_model.Sub_id+1)
	return_col := fmt.Sprintf("return_%d", timer_model.Sub_id+1)

	// Prepare the SQL statement using placeholders for parameters
	sql_statement := fmt.Sprintf(`
        INSERT INTO SubmersibleTimers (fc, %s, %s)
        VALUES (?, ?, ?)
        ON CONFLICT (fc)
        DO UPDATE SET fc = ?, %s = ?, %s = ?;
    `, sub_col, return_col, sub_col, return_col)

	// Prepare the statement
	statement, err := db.Prepare(sql_statement)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(timer_model.Fc_id, timer_model.Sub_id, timer_model.Return_time)
	if err != nil {
		log.Fatal(err)
	}
}
