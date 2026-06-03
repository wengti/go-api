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
		router.POST("/events", addEvent)
		router.PUT("events/:id", updateEventById)
		router.DELETE("events/:id", deleteEventById)
	}

	router.POST("/signup", handleSignup)
	router.POST("/login", handleLogin)
}
