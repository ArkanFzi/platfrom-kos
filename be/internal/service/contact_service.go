package service

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type ContactService interface {
	SendContactMessage(name, email, message string) error
}

type contactService struct {
	smtpHost     string
	smtpPort     int
	smtpEmail    string
	smtpPassword string
	targetEmail  string
}

func NewContactService() ContactService {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if port == 0 {
		port = 587 // default SMTP port
	}

	return &contactService{
		smtpHost:     os.Getenv("SMTP_HOST"),
		smtpPort:     port,
		smtpEmail:    os.Getenv("SMTP_EMAIL"),
		smtpPassword: os.Getenv("SMTP_PASSWORD"),
		targetEmail:  os.Getenv("CONTACT_EMAIL"),
	}
}

func (s *contactService) SendContactMessage(name, email, message string) error {
	// Create new email message
	m := gomail.NewMessage()
	
	// Set email headers
	m.SetHeader("From", s.smtpEmail)
	m.SetHeader("To", s.targetEmail)
	m.SetHeader("Subject", fmt.Sprintf("Pesan Baru dari %s - Contact Form Koskosan", name))
	m.SetHeader("Reply-To", email)
	
	// Create HTML email body
	htmlBody := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<style>
		body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f4f4f4; margin: 0; padding: 0; }
		.container { max-width: 600px; margin: 40px auto; background: white; border-radius: 12px; overflow: hidden; box-shadow: 0 4px 6px rgba(0,0,0,0.1); }
		.header { background: linear-gradient(135deg, #292524 0%%, #44403c 100%%); color: white; padding: 30px; text-align: center; }
		.header h1 { margin: 0; font-size: 24px; font-weight: 700; }
		.header p { margin: 10px 0 0 0; opacity: 0.9; font-size: 14px; }
		.content { padding: 40px 30px; }
		.info-row { margin-bottom: 20px; }
		.info-row label { display: block; color: #78716c; font-size: 12px; font-weight: 600; text-transform: uppercase; letter-spacing: 0.5px; margin-bottom: 5px; }
		.info-row .value { color: #292524; font-size: 16px; font-weight: 500; }
		.message-box { background-color: #fafaf9; border-left: 4px solid #78716c; padding: 20px; border-radius: 8px; margin-top: 10px; }
		.message-box p { margin: 0; color: #292524; line-height: 1.6; white-space: pre-wrap; }
		.footer { background-color: #fafaf9; padding: 20px; text-align: center; border-top: 1px solid #e7e5e4; }
		.footer p { margin: 5px 0; color: #78716c; font-size: 13px; }
	</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h1>📬 Pesan Baru dari Contact Form</h1>
			<p>Koskosan Rahmat ZAW - Malang</p>
		</div>
		<div class="content">
			<div class="info-row">
				<label>Dari</label>
				<div class="value">%s</div>
			</div>
			<div class="info-row">
				<label>Email</label>
				<div class="value">%s</div>
			</div>
			<div class="info-row">
				<label>Pesan</label>
				<div class="message-box">
					<p>%s</p>
				</div>
			</div>
		</div>
		<div class="footer">
			<p>Email ini dikirim secara otomatis dari website Koskosan Rahmat ZAW</p>
			<p>Untuk membalas, klik "Reply" atau email ke: %s</p>
		</div>
	</div>
</body>
</html>
	`, name, email, message, email)
	
	// Set HTML body
	m.SetBody("text/html", htmlBody)
	
	// Create plain text alternative
	plainBody := fmt.Sprintf(`
Pesan Baru dari Contact Form
Koskosan Rahmat ZAW - Malang

Dari: %s
Email: %s

Pesan:
%s

---
Email ini dikirim secara otomatis dari website Koskosan Rahmat ZAW
Untuk membalas, email ke: %s
	`, name, email, message, email)
	
	m.AddAlternative("text/plain", plainBody)
	
	// Create SMTP dialer
	d := gomail.NewDialer(s.smtpHost, s.smtpPort, s.smtpEmail, s.smtpPassword)
	
	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	
	return nil
}
