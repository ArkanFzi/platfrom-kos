package utils

import (
	"fmt"
	"koskosan-be/internal/config"
	"log"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

type EmailSender interface {
	SendResetPasswordEmail(toEmail, token string) error
	SendPaymentSuccessEmail(toEmail, tenantName string, amount float64, date time.Time) error
	SendPaymentReminderEmail(toEmail, tenantName string, amount float64, dueDate time.Time, paymentLink string) error
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

func (s *GomailSender) SendPaymentSuccessEmail(toEmail, tenantName string, amount float64, date time.Time) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Payment Confirmation - Kost Putra Rahmat ZAW")

	formattedAmount := fmt.Sprintf("Rp %.0f", amount)
	formattedDate := date.Format("02 Jan 2006 15:04")

	body := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; padding: 20px; color: #333;">
			<h2 style="color: #4CAF50;">Pembayaran Berhasil!</h2>
			<p>Halo, <strong>%s</strong>,</p>
			<p>Terima kasih, pembayaran Anda telah kami terima.</p>
			<table style="width: 100%%; max-width: 400px; margin: 20px 0; border-collapse: collapse;">
				<tr>
					<td style="padding: 8px; border-bottom: 1px solid #ddd;">Jumlah</td>
					<td style="padding: 8px; border-bottom: 1px solid #ddd; font-weight: bold;">%s</td>
				</tr>
				<tr>
					<td style="padding: 8px; border-bottom: 1px solid #ddd;">Tanggal</td>
					<td style="padding: 8px; border-bottom: 1px solid #ddd;">%s</td>
				</tr>
				<tr>
					<td style="padding: 8px; border-bottom: 1px solid #ddd;">Status</td>
					<td style="padding: 8px; border-bottom: 1px solid #ddd; color: green;">Lunas</td>
				</tr>
			</table>
			<p>Simpan email ini sebagai bukti pembayaran yang sah.</p>
			<br/>
			<p style="font-size: 12px; color: #888;">Kost Putra Rahmat ZAW Management</p>
		</div>
	`, tenantName, formattedAmount, formattedDate)

	m.SetBody("text/html", body)
	return s.dialer.DialAndSend(m)
}

func (s *GomailSender) SendPaymentReminderEmail(toEmail, tenantName string, amount float64, dueDate time.Time, paymentLink string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Tagihan Belum Dibayar - Kost Putra Rahmat ZAW")

	formattedAmount := fmt.Sprintf("Rp %.0f", amount)
	formattedDate := dueDate.Format("02 Jan 2006")

	body := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; padding: 20px; color: #333;">
			<h2 style="color: #F44336;">Pengingat Tagihan</h2>
			<p>Halo, <strong>%s</strong>,</p>
			<p>Ini adalah pengingat untuk tagihan sewa kos Anda yang belum dibayar.</p>
			<table style="width: 100%%; max-width: 400px; margin: 20px 0; border-collapse: collapse;">
				<tr>
					<td style="padding: 8px; border-bottom: 1px solid #ddd;">Total Tagihan</td>
					<td style="padding: 8px; border-bottom: 1px solid #ddd; font-weight: bold;">%s</td>
				</tr>
				<tr>
					<td style="padding: 8px; border-bottom: 1px solid #ddd;">Jatuh Tempo</td>
					<td style="padding: 8px; border-bottom: 1px solid #ddd; color: #F44336;">%s</td>
				</tr>
			</table>
			<p>Mohon segera lakukan pembayaran sebelum tanggal jatuh tempo untuk menghindari denda.</p>
			<div style="margin: 30px 0;">
				<a href="%s" style="background-color: #2196F3; color: white; padding: 12px 24px; text-decoration: none; border-radius: 4px; font-weight: bold;">Bayar Sekarang</a>
			</div>
			<p style="font-size: 12px; color: #888;">Jika Anda sudah melakukan pembayaran, mohon abaikan email ini.</p>
		</div>
	`, tenantName, formattedAmount, formattedDate, paymentLink)

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

func (s *LogSender) SendPaymentSuccessEmail(toEmail, tenantName string, amount float64, date time.Time) error {
	log.Printf("---------------------------------------------------------")
	log.Printf("[EMAIL SIMULATION] To: %s", toEmail)
	log.Printf("[EMAIL SIMULATION] Subject: Payment Success")
	log.Printf("[EMAIL SIMULATION] Body: Payment of Rp %.0f received on %s", amount, date.Format("02 Jan 2006"))
	log.Printf("---------------------------------------------------------")
	return nil
}

func (s *LogSender) SendPaymentReminderEmail(toEmail, tenantName string, amount float64, dueDate time.Time, paymentLink string) error {
	log.Printf("---------------------------------------------------------")
	log.Printf("[EMAIL SIMULATION] To: %s", toEmail)
	log.Printf("[EMAIL SIMULATION] Subject: Payment Reminder")
	log.Printf("[EMAIL SIMULATION] Body: Please pay Rp %.0f by %s. Link: %s", amount, dueDate.Format("02 Jan 2006"), paymentLink)
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
