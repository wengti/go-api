package db

import "example.com/go-api/models"

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
