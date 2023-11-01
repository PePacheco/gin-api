package middlewares

import (
	"net/http"
	"notes-app/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		splits := strings.Split(authHeader, " ")

		if len(splits) != 2 || splits[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token format is Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := splits[1]

		if !helpers.ValidateToken(tokenString) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
