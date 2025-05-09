package main

import (
	"crud-service/internal/db"
	"crud-service/internal/route"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	err := db.InitDB() // Calls InitDB from the db package
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB() // Make sure to close the DB when main ends

	// Create a new Gin router
	router := gin.Default()

	// Register routes
	route.RegisterRoutes(router)
	route.LoginRoutes(router)

	// Start the Gin server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
