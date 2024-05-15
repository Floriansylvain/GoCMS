package gateways

import (
	"GoCMS/domain/gateways"
	"bytes"
	"embed"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"strconv"
)

type MailRepository struct{}

//go:embed mail/templates/*
var mailTemplateFiles embed.FS

func (m MailRepository) Send(receiverAddress string, templateName string, data interface{}) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	tmpl, err := template.ParseFS(mailTemplateFiles, "mail/templates/"+templateName+".html")
	if err != nil {
		log.Println("error when template.ParseFS:", err)
		return err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		log.Println("error when tmpl.Execute:", err)
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeaders(map[string][]string{
		"From":         {"GoCMS <" + from + ">"},
		"To":           {receiverAddress},
		"MIME-version": {"1.0"},
		"Content-Type": {"text/html"},
		"charset":      {"UTF-8"},
		"Subject":      {"GoCMS | Action required"},
	})
	msg.SetBody("text/html", body.String())

	if err := d.DialAndSend(msg); err != nil {
		log.Println("error when sending email:", err)
		return err
	}

	return nil
}

var _ gateways.IMailRepository = &MailRepository{}
