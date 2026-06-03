package db

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
