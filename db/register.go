package db

import "errors"

func RegisterEvent(eventId, userId int64) error {
	// Prepare query statement
	query := `
	INSERT INTO registration (event_id, user_id)
	VALUES ($1, $2)
	`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query statement
	_, err = stmt.Exec(eventId, userId)
	if err != nil {
		return err
	}

	// Return no error
	return nil
}

func DeregisterEvent(eventId, userId int64) error {
	// Prepare query statement
	query := `
	DELETE FROM registration
	WHERE event_id = $1 and user_id = $2
	`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query statement
	result, err := stmt.Exec(eventId, userId)
	if err != nil {
		return err
	}

	row_affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if row_affected == 0 {
		return errors.New("No deregistration is made.")
	}

	// Return no error
	return nil
}
