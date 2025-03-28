package middleware

import (
	"net/http"
	"pricing-app/config"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware sets up Cross-Origin Resource Sharing (CORS) headers for HTTP requests.
// It allows requests from the frontend origin specified in the application configuration.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Envs.CORSFrontendOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
