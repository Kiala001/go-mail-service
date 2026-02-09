package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kiala001/go-mail-service/internal/apikey"
)

func GenerateKey(c *gin.Context) {
	key, err := apikey.Generate()
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao gerar API Key"})
		return
	}
	c.JSON(200, gin.H{"api_key": key})
}