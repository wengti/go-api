package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitializeDatabase() {

	// Create database sql file
	filename := "api.sql"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsExist(err) {
			// file already exists, skip
		} else {
			panic("Fail to create the database file.")
		}
	}
	defer file.Close()

	// Connect to the database
	Db, err = sql.Open("sqlite3", filename)
	if err != nil {
		panic(err.Error())
	}

	// Create User Table if not exists
	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE
	)
	`
	_, err = Db.Exec(createUserTableQuery)
	if err != nil {
		panic(err.Error())
	}

	// Create event table if not exists
	createEventTableQuery := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		location TEXT NOT NULL,
		description TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	_, err = Db.Exec(createEventTableQuery)
	if err != nil {
		panic(err.Error())
	}

}
