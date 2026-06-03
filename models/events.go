package models

type Event struct {
	Id          int64
	Name        string `binding:"required"`
	Location    string `binding:"required"`
	Description string `binding:"required"`
	UserId      int64
}
