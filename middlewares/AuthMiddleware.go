package middlewares

import (
	"fmt"
	"net/http"
	"notes-app/helpers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		splits := strings.Split(authHeader, " ")

		// Check if header format is correct
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

func validateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return helpers.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return false
	}
	return true
}
