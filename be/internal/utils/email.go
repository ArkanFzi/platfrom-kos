package utils

import (
	"fmt"
	"koskosan-be/internal/config"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailSender interface {
	SendResetPasswordEmail(toEmail, token string) error
}

type GomailSender struct {
	dialer *gomail.Dialer
	from   string
}

func NewGomailSender(cfg *config.Config) *GomailSender {
	port, _ := strconv.Atoi(cfg.SMTPPort)
	dialer := gomail.NewDialer(cfg.SMTPHost, port, cfg.SMTPEmail, cfg.SMTPPassword)
	return &GomailSender{
		dialer: dialer,
		from:   cfg.SMTPEmail,
	}
}

func (s *GomailSender) SendResetPasswordEmail(toEmail, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Reset Password Request")
	
	// Better to use a configurable frontend URL, but for now hardcode or assume localhost/production
	// TODO: Add frontend URL to config
	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", token)

	body := fmt.Sprintf(`
		<h3>Reset Password Request</h3>
		<p>You requested a password reset. Click the link below to reset your password:</p>
		<p><a href="%s">Reset Password</a></p>
		<p>If you did not request this, please ignore this email.</p>
	`, resetLink)

	m.SetBody("text/html", body)

	return s.dialer.DialAndSend(m)
}

type LogSender struct{}

func (s *LogSender) SendResetPasswordEmail(toEmail, token string) error {
	log.Printf("---------------------------------------------------------")
	log.Printf("[EMAIL SIMULATION] To: %s", toEmail)
	log.Printf("[EMAIL SIMULATION] Subject: Reset Password Request")
	log.Printf("[EMAIL SIMULATION] Token: %s", token)
	log.Printf("[EMAIL SIMULATION] Link: http://localhost:3000/reset-password?token=%s", token)
	log.Printf("---------------------------------------------------------")
	return nil
}

// Helper to choose sender
func NewEmailSender(cfg *config.Config) EmailSender {
	if cfg.SMTPHost != "" && cfg.SMTPEmail != "" && cfg.SMTPPassword != "" {
		return NewGomailSender(cfg)
	}
	log.Println("SMTP credentials not found, using LogSender (Simulation Mode)")
	return &LogSender{}
}
