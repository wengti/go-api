package routes

import (
	"example.com/go-api/middlewares"
	"github.com/gin-gonic/gin"
)

func HandleRoutes(router *gin.Engine) {

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventById)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.VerifyAuthorization())
	{
		authGroup.POST("/events", addEvent)
		authGroup.PUT("events/:id", updateEventById)
		authGroup.DELETE("events/:id", deleteEventById)
	}

	router.POST("/signup", handleSignup)
	router.POST("/login", handleLogin)
}
