package db

import (
	"database/sql"
	"os"

	"example.com/go-api/models"
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

func SaveNewEvent(event *models.Event) error {

	// Prepare query statement
	query := `
	INSERT INTO events (name, location, description, user_id)
	VALUES ($1, $2, $3, $4)
	`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute statement
	_, err = stmt.Exec(event.Name, event.Location, event.Description, event.UserId)
	if err != nil {
		return err
	}

	return nil

}

func FetchAllEvent() ([]models.Event, error) {

	// Placeholder
	fetchedEvents := []models.Event{}

	// Execute query statement
	query := `
	SELECT * FROM events
	`
	rows, err := Db.Query(query)
	if err != nil {
		return []models.Event{}, err
	}
	defer rows.Close()

	// Save the return results into placeholder
	for rows.Next() {
		var name, location, description string
		var id, userId int64
		err := rows.Scan(&id, &name, &location, &description, &userId)
		if err != nil {
			return []models.Event{}, err
		}
		fetchedEvents = append(fetchedEvents, models.Event{
			Id:          id,
			Name:        name,
			Location:    location,
			Description: description,
			UserId:      userId,
		})
	}

	// Return results when no error
	return fetchedEvents, nil

}

func FetchEventById(targetId int64) (models.Event, error) {
	// Placeholder
	var name, location, description string
	var id, userId int64

	// Execute query statement
	query := `
	SELECT * FROM events
	WHERE id = $1
	`
	row := Db.QueryRow(query, targetId)
	err := row.Scan(&id, &name, &location, &description, &userId)
	if err != nil {
		return models.Event{}, err
	}

	return models.Event{
		Id:          id,
		Name:        name,
		Location:    location,
		Description: description,
		UserId:      userId,
	}, nil
}
