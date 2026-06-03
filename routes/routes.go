package routes

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.Engine) {

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventById)
	router.POST("/events", addEvent)
}
