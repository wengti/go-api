package db

import "example.com/go-api/models"

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
