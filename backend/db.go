package backend

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	// connect database
	db, connectErr := sql.Open("sqlite3", "./sqlite.db")
	if connectErr != nil {
		fmt.Println("Counld not open database: %+v", connectErr)
		panic(connectErr)
	}

	// create table
	_, execErr := db.Exec(
		`CREATE TABLE IF NOT EXISTS "TASKS" ("ID" INTEGER PRIMARY KEY, "TITLE" STRING)`,
	)
	if execErr != nil {
		fmt.Println("Could not create table: %+v", execErr)
		panic(execErr)
	}
}
