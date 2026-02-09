package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kiala001/go-mail-service/internal/apikey"
)

func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API Key ausente"})
			return
		}

		key := strings.TrimPrefix(auth, "Bearer ")
		if !apikey.IsValid(key) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API Key inválida"})
			return
		}
		c.Next()
	}
}