package backend

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func InitDB() {
	// open database
	db, openErr := OpenDB()
	if openErr != nil {
		panic(openErr)
	}

	// create table
	_, execErr := db.Exec(
		`CREATE TABLE IF NOT EXISTS "TASKS" (` +
			`"ID" INTEGER PRIMARY KEY, ` +
			`"NAME" STRING,` +
			`"MEMO" STRING,` +
			`"QUADRANT" INT,` +
			`"COMPLETEFLAG" BOOL,` +
			`"CREATEDAT" DATETIME,` +
			`"UPDATEDAT" DATETIME` +
			`)`,
	)
	if execErr != nil {
		log.Printf("Could not create table: %+v", execErr)
		panic(execErr)
	}
}

func OpenDB() (*sql.DB, error) {
	var dbFile string
	if os.Getenv("RUNNING_ENV") != "TEST" {
		dbFile = "./sqlite3.db"
	} else {
		dbFile = "./test.db"
	}
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Printf("Unexpected error occured during open database: %+v", err)
		return nil, err
	}
	return db, nil
}
