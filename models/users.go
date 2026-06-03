package models

type User struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
