// routes.go
package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	// Call setup functions for user and post routes
	SetupHelloRoutes(r)
	// Add more setup functions for other routes if needed
}
