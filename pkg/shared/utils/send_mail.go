package utils

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

// DataVerifyUser struct
type DataVerifyUser struct {
	Email string
	Token string
}

// SendMail func
func SendMail(email string, templateSendMail string, data interface{}) error {
	var message string

	workDir := GetStringFlag("workdir")
	var (
		basePath     = filepath.Join(workDir, "/web/mail")
		templatePath = filepath.Join(basePath, "template")
		indexFile    = filepath.Join(templatePath, templateSendMail)
	)
	t, err := template.New(templateSendMail).ParseFiles(indexFile)
	if err != nil {
		return err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return err
	}

	message = tpl.String()

	username := os.Getenv("SMTP_USER")
	smtpEmail := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	host := os.Getenv("SMTP_HOST")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))

	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", smtpEmail)
	m.SetHeader("To", email)
	m.SetBody("text/html", message)

	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
