package routes

import (
	"net/http"

	"example.com/go-api/db"
	"example.com/go-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func handleSignup(context *gin.Context) {

	// Get user email and pw from request body
	newUser := models.User{}
	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Hash password
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	newUser.Password = string(hashedPw)

	// Save user data into database
	err = db.SaveNewUser(newUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully sign up!"})

}
