package middlewares

import (
	"crud-app/models"
	"crud-app/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func BloggerRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, err := utils.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		if claims["role"] != models.BloggerRole {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized blogger"})
			c.Abort()
			return
		}
		c.Next()
	}
}
