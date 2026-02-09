package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kiala001/go-mail-service/internal/config"
	"github.com/kiala001/go-mail-service/internal/handlers"
	"github.com/kiala001/go-mail-service/internal/middleware"
	"github.com/kiala001/go-mail-service/internal/services"
	"github.com/kiala001/go-mail-service/pkg/mailer"
)

func main() {
	r := gin.Default()

	m := mailer.NewMailer()
	svc := services.NewEmailService(m)

	api := r.Group("/api")
	{
		api.POST("/generate-key", handlers.GenerateKey)
		api.Use(middleware.ApiKeyAuth())
		api.POST("/send-email", handlers.SendEmail(svc))
		api.POST("/send-template", handlers.SendTemplate(svc))
	}

	r.Run(":" + config.AppPort)
}