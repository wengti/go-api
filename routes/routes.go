package routes

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.Engine) {

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventById)
	router.POST("/events", addEvent)
	router.PUT("events/:id", updateEventById)
	router.DELETE("events/:id", deleteEventById)

	router.POST("/signup", handleSignup)
	router.POST("/login", handleLogin)
}
