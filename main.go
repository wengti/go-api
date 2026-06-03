package main

import (
	"example.com/go-api/db"
	"example.com/go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDatabase()
	router := gin.Default()
	routes.HandleRoutes(router)
	router.Run()
}
