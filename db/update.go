package db

import "example.com/go-api/models"

func UpdateEventById(updatedEvent models.Event) error {

	// Prepare query statement
	query := `
	UPDATE events
	SET name = $1, location = $2, description = $3
	WHERE id = $4
	`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query statement
	_, err = Db.Exec(query, updatedEvent.Name, updatedEvent.Location, updatedEvent.Description, updatedEvent.Id)
	if err != nil {
		return err
	}

	return nil

}
