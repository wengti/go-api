package main

import (
	"example.com/go-api/db"
	"example.com/go-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	db.InitializeDatabase()
	router := gin.Default()
	routes.HandleRoutes(router)
	router.Run()
}
