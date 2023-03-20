package commons

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/mail"
	"os"
	"strconv"
	"text/template"

	"crypto/sha256"

	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/skip2/go-qrcode"
	"gopkg.in/gomail.v2"
)

func ParseEmail(address string) (string, error) {
	if email, err := mail.ParseAddress(address); err != nil {
		return "", err
	} else {
		return email.Address, nil
	}
}

func SendEmail(to string, plainContent *string, templatePath *string, data any, subject *string, attachment *string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", env.Env.EMAIL_FROM)
	m.SetHeader("To", to)
	if plainContent != nil {
		m.SetBody("text/plain", *plainContent)
	} else {
		var body bytes.Buffer
		bodyTemplate, err := template.ParseFiles(*templatePath)
		if err != nil {
			return err
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
	// fmt.Println(d)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func Hash(data string) string {
	hashedByte := sha256.Sum256([]byte(data))
	hashedString := fmt.Sprintf("%v", hashedByte[:])
	return hashedString
}

func CreateQrCode(text, path string) error {
	return qrcode.WriteFile(text, qrcode.Medium, 256, path)
}

func SaveFileFromBinary(base64Str, toPath string) error {
	dec, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}
	file, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(dec); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	return nil
}
