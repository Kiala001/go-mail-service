package services

import (
	"github.com/kiala001/go-mail-service/internal/templates"
	"github.com/kiala001/go-mail-service/pkg/mailer"
)

type EmailService struct {
	Mailer *mailer.Mailer
}

func NewEmailService(m *mailer.Mailer) *EmailService {
	return &EmailService{Mailer: m}
}

func (s *EmailService) SendHTML(from string, to []string, subject, html string) error {
	return s.Mailer.Send(from, to, subject, html)
}

func (s *EmailService) SendTemplate(from string, to []string, subject, tpl string, data any) error {
	html, err := templates.Render(tpl, data)
	if err != nil {
		return err
	}
	return s.SendHTML(from, to, subject, html)
}