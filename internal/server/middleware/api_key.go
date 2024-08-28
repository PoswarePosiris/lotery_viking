package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

// CheckAPIKey checks if the request contains a valid API key
func CheckAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("api-key")
		expectedAPIKey := os.Getenv("API_KEY")

		if apiKey != expectedAPIKey {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden access"})
			return
		}

		c.Next()
	}
}
