package cmd

import "github.com/gin-gonic/gin"

func App() {
	// Initialize Gin router
	router := gin.Default()

	router.Run(":8000")
}
