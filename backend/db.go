package backend

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
		fmt.Println("Could not create table: %+v", execErr)
		panic(execErr)
	}
}

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		fmt.Println("Unexpected error occured during open database: %+v", err)
		return nil, err
	}
	return db, nil
}
