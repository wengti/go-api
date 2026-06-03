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

func FetchUserByEmail(email string) (int64, string, error) {

	var userId int64
	var hashedPw string

	// Prepare query statement
	query := `
	SELECT id, password FROM users
	WHERE email = $1
	`

	// Execute query statement
	row := Db.QueryRow(query, email)
	err := row.Scan(&userId, &hashedPw)
	if err != nil {
		return -1, "", err
	}

	return userId, hashedPw, nil

}
