package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckKiosk() gin.HandlerFunc {
	return func(c *gin.Context) {
		macKiosk := c.GetHeader("Authorization")
		if macKiosk != "" {
			macKiosk = strings.TrimPrefix(macKiosk, "Bearer ")
			ctx := context.WithValue(c.Request.Context(), "macKiosk", macKiosk)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			c.AbortWithStatusJSON(403, gin.H{"error": "Invalid Kiosk ID"})
		}
	}
}
