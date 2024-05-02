// user_routes.go
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupHelloRoutes(r *gin.Engine) {
	// Define user routes here
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	// Add more routes as needed
}
