package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"embed"
)

//go:embed templates/*.html
var templateFS embed.FS

func Send(to, subject, body string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	from := os.Getenv("SMTP_FROM")
	fromName := os.Getenv("SMTP_FROM_NAME")
	if fromName == "" {
		fromName = "The Nostos Team" // fallback name if env var is not set
	}

	log.Printf("SMTP Config - Host: %s, Port: %s, From: %s", host, port, from)

	if host == "" || port == "" || from == "" {
		return fmt.Errorf("SMTP configuration is incomplete. Please check environment variables")
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	msg := fmt.Sprintf("From: %s <%s>\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		fromName, from, to, subject, body)

	// Create custom client with TLS verification disabled
	c, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer c.Close()

	if err = c.Mail(from); err != nil {
		return fmt.Errorf("MAIL FROM error: %v", err)
	}

	if err = c.Rcpt(to); err != nil {
		return fmt.Errorf("RCPT TO error: %v", err)
	}

	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("DATA error: %v", err)
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		return fmt.Errorf("body write error: %v", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("close error: %v", err)
	}

	return nil
}

func SendConfirmationEmail(to, name, link string) error {
	body, err := RenderTemplate("templates/welcomeMail.html", map[string]string{
		"Name": name,
		"Link": link,
	})
	if err != nil {
		return err
	}
	return Send(to, "Welcome to the Nostos family", body)
}

func SendPasswordResetEmail(to, link string) error {
	body, err := RenderTemplate("templates/resetPassword.html", map[string]string{
		"Link": link,
	})
	if err != nil {
		return err
	}
	return Send(to, "Reset your password", body)
}

func RenderTemplate(templatePath string, data any) (string, error) {
	tmpl, err := template.ParseFS(templateFS, templatePath)
	if err != nil {
		log.Printf("❌ Error parsing template: %v", err)
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Printf("❌ Error executing template: %v", err)
		return "", err
	}

	return buf.String(), nil
}
