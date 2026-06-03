package db

import "example.com/go-api/models"

// Save

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

// Fetch

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

// Update

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

// Delete

func DeleteEventById(targetId int64) error {

	// Prepare query statement
	query := `
	DELETE FROM events
	WHERE id = $1
	`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute statement
	_, err = stmt.Exec(targetId)
	if err != nil {
		return err
	}

	return nil
}
