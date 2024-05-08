package db

import (
	"LootData4/lodestone"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Update_FC_Name(id string, world string) string {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM FC_World WHERE fc = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	var updated int64 = 0
	var name string = ""

	for rows.Next() {
		var _fc string
		var _world string
		err = rows.Scan(&_fc, &_world, &name, &updated)
		if err != nil {
			log.Fatal(err)
		}
	}
	rows.Close()

	if time.Now().Unix()-updated < 604800 {
		//return if it hasnt been a week
		return name
	}

	name, err = lodestone.Get_FC_Name(id)
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare("REPLACE INTO FC_World (fc, world, name, updated) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(id, world, name, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}

	return name
}
