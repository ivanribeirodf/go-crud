package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue, exists := c.Get("token")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não encontrado"})
			c.Abort()
			return
		}

		token := tokenValue.(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permissão negada"})
			c.Abort()
			return
		}

		c.Next()
	}
}
