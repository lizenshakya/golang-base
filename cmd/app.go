package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lizenshakya/golang-base/config"
	"github.com/lizenshakya/golang-base/internal/app/database"
	"github.com/lizenshakya/golang-base/internal/app/routes"
)

func init() {
	config.LoadEnvVariables()
	database.ConnectToDb()
}

func App() {
	// Initialize Gin router
	router := gin.Default()
	routes.SetupRoutes(router)
	if err := router.Run(":8000"); err != nil {
		// If an error occurs, log it and exit the program
		log.Fatalf("Error running server: %v", err)
	}
}
