package db

import "example.com/go-api/models"

func SaveNewUser(newUser models.User) error {

	// Prepare query statement
	query := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query statement
	_, err = stmt.Exec(newUser.Email, newUser.Password)
	if err != nil {
		return err
	}

	return nil
}
