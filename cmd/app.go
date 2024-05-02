package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/lizenshakya/golang-base/internal/app/routes"
)

func App() {
	// Initialize Gin router
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8000")
}
