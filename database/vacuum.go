package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Vacuum() {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("VACUUM")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database vacuumed successfully")
}
