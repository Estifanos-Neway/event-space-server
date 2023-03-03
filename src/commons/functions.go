package commons

import (
	"bytes"
	"net/mail"
	"strconv"
	"text/template"

	"github.com/estifanos-neway/event-space-server/src/env"
	"gopkg.in/gomail.v2"
)

func ParseEmail(address string) (string, error) {
	if email, err := mail.ParseAddress(address); err != nil {
		return "", err
	} else {
		return email.Address, nil
	}
}

func SendEmail(to string, plainContent *string, templatePath *string, data *any, subject *string, attachment *string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", env.Env.EMAIL_FROM)
	m.SetHeader("To", to)
	if plainContent != nil {
		m.SetBody("text/plain", *plainContent)
	} else {
		var body bytes.Buffer
		bodyTemplate, err := template.ParseFiles(*templatePath)
		if err != nil {
			return nil
		}
		if err := bodyTemplate.Execute(&body, data); err != nil {
			return err
		}
		m.SetBody("text/html", body.String())
	}
	if subject != nil {
		m.SetHeader("Subject", *subject)
	}
	if attachment != nil {
		m.Attach(*attachment)
	}
	port, err := strconv.ParseInt(env.Env.SMTP_PORT, 10, 64)
	if err != nil {
		return err
	}
	d := gomail.NewDialer(env.Env.SMTP_HOST, int(port), env.Env.SMTP_USERNAME, env.Env.SMTP_PASSWORD)

	if err := d.DialAndSend(m); err != nil {
		return nil
	}
	return nil
}
