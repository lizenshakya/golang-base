package cmd

import (
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
	router.Run(":8000")
}
