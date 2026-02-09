package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kiala001/go-mail-service/internal/services"
)

type EmailRequest struct {
	To       []string          `json:"to"`
	Subject  string            `json:"subject"`
	Html     string            `json:"html"`
	Template string            `json:"template"`
	Data     map[string]any    `json:"data"`
	From     string            `json:"from"`
}

func SendEmail(s *services.EmailService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req EmailRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "JSON inválido"})
			return
		}

		err := s.SendHTML(req.From, req.To, req.Subject, req.Html)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "Email enviado"})
	}
}

func SendTemplate(s *services.EmailService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req EmailRequest
		c.ShouldBindJSON(&req)

		err := s.SendTemplate(req.From, req.To, req.Subject, req.Template, req.Data)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "Email enviado"})
	}
}