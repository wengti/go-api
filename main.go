package main

import (
	"net/http"
	"strconv"

	"example.com/go-api/db"
	"example.com/go-api/models"
	"github.com/gin-gonic/gin"
)

// var allEvents = []models.Event{
// 	{
// 		Id:          1,
// 		Name:        "Event 1",
// 		Location:    "Venue 1",
// 		Description: "A test event",
// 		UserId:      1,
// 	},
// 	{
// 		Id:          2,
// 		Name:        "Event 2",
// 		Location:    "Venue 2",
// 		Description: "A test event",
// 		UserId:      2,
// 	},
// }

func main() {
	db.InitializeDatabase()
	router := gin.Default()

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventById)
	router.POST("/events", addEvent)

	router.Run()
}

// Handlers

func getEvents(context *gin.Context) {
	allEvents, err := db.FetchAllEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": allEvents})
}

func getEventById(context *gin.Context) {

	// Get event ID from path and convert to integer
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Find the corresponding target event
	targetEvent, err := db.FetchEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return the found target event
	context.JSON(http.StatusOK, gin.H{"events": targetEvent})
}

func addEvent(context *gin.Context) {

	// Get the body request and bind it to the event struct
	newEvent := models.Event{}
	err := context.ShouldBindJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Add the user id to the new event since not included in body request
	newEvent.UserId = 1

	// Add the new event to the lists
	err = db.SaveNewEvent(&newEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"messsage": "Event is added successfully."})
}
