package routes

import (
	"net/http"
	"os"

	"example.com/go-api/db"
	"example.com/go-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func handleLogin(context *gin.Context) {

	// Get user email and pw from request body
	loginUser := models.User{}
	err := context.ShouldBindJSON(&loginUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Get hashed pw from database
	userId, hashedPw, err := db.FetchUserByEmail(loginUser.Email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credential."})
		return
	} else if userId == -1 || hashedPw == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credential."})
		return
	}

	// Compare password and hashedPw
	err = bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(loginUser.Password))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credential."})
		return
	}

	// Get a signed JWT string when password matches
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    userId,
			"email": loginUser.Email,
		})
	jwtString, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return signed JWT string
	context.JSON(http.StatusOK, gin.H{"jwt_string": jwtString})
}
