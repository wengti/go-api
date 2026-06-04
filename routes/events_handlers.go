package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/go-api/db"
	"example.com/go-api/models"
	"github.com/gin-gonic/gin"
)

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
	// jwt token stores number as flt64, text as string
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	fltUserId, ok := userId.(float64)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	newEvent.UserId = int64(fltUserId)

	// Add the new event to the lists
	err = db.SaveNewEvent(&newEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"messsage": "Event is added successfully."})
}

func updateEventById(context *gin.Context) {

	// Get event ID from path and convert to integer
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Get the body request and bind it to the event struct
	updatedEvent := models.Event{}
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Bind eventId into the updatedEvent
	updatedEvent.Id = eventId

	// Add the user id to the new event since not included in body request
	// jwt token stores number as flt64, text as string
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	fltUserId, ok := userId.(float64)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// Bind userId into the updatedEvent
	updatedEvent.UserId = int64(fltUserId)

	// Hit database to update the event
	err = db.UpdateEventById(updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Event with id %v is updated successfully.", eventId)})
}

func deleteEventById(context *gin.Context) {
	// Get event ID from path and convert to integer
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Add the user id to the new event since not included in body request
	// jwt token stores number as flt64, text as string
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	fltUserId, ok := userId.(float64)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	err = db.DeleteEventById(eventId, int64(fltUserId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Event with id %v is deleted successfully.", eventId)})
}
