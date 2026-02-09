package mailer

import (
	"strconv"

	"github.com/kiala001/go-mail-service/internal/config"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Dialer *gomail.Dialer
}

func NewMailer() *Mailer {
	port, _ := strconv.Atoi(config.SMTPPort)

	d := gomail.NewDialer(
		config.SMTPHost,
		port,
		config.SMTPUser,
		config.SMTPPass,
	)

	return &Mailer{Dialer: d}
}

func (m *Mailer) Send(from string, to []string, subject, html string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", html)

	return m.Dialer.DialAndSend(msg)
}