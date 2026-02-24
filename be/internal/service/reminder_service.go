package service

import (
	"fmt"
	"koskosan-be/internal/models"
	"koskosan-be/internal/repository"
	"koskosan-be/internal/utils"
	"time"

	"gorm.io/gorm"
)

type ReminderService interface {
	CreateMonthlyReminders() error
	SendPendingReminders() ([]models.PaymentReminder, error)
	MarkReminderAsSent(reminderID uint) error
	MarkPaymentAsPaid(reminderID uint) error
	GetPendingReminders() ([]models.PaymentReminder, error)
}

type reminderService struct {
	paymentRepo repository.PaymentRepository
	db          *gorm.DB
	emailSender utils.EmailSender
	waSender    utils.WhatsAppSender
}

func NewReminderService(paymentRepo repository.PaymentRepository, db *gorm.DB, emailSender utils.EmailSender, waSender utils.WhatsAppSender) ReminderService {
	return &reminderService{paymentRepo, db, emailSender, waSender}
}

// CreateMonthlyReminders membuat reminder untuk semua pembayaran DP yang jatuh tempo
func (s *reminderService) CreateMonthlyReminders() error {
	// Ambil semua pembayaran dengan tipe "dp" yang belum lunas (belum Settled)
	var payments []models.Pembayaran
	// Asumsi logic: pembayaran DP yang sudah Confirmed tapi belum lunas keseluruhannya (masih nyicil/bulanan)
	// Kita cari pembayaran yang TipePembayaran = 'dp' dan StatusPembayaran = 'Confirmed'
	// Note: Logic bisnis ini mungkin perlu disesuaikan dengan flow 'Settled' vs 'Confirmed'
	if err := s.db.Where("tipe_pembayaran = ? AND status_pembayaran = ?", "dp", "Confirmed").Find(&payments).Error; err != nil {
		return err
	}

	for _, p := range payments {
		// Cek apakah sudah ada reminder untuk bulan ini/periode ini
		// Logic sederhana: Reminder dibuat jika belum ada reminder yang 'Pending' atau 'Sent' untuk pembayaran ini yang tanggalnya > hari ini
		// Atau lebih baik: Cek apakah TanggalJatuhTempo sudah dekat

		// Misalnya: Jatuh tempo setiap tanggal X. Kita buat reminder H-3.
		// Disini kita simulasikan pembuatan reminder jika belum ada reminder pending.
		
		var count int64
		s.db.Model(&models.PaymentReminder{}).
			Where("pembayaran_id = ? AND status_reminder IN ?", p.ID, []string{"Pending", "Sent"}).
			Count(&count)

		if count == 0 {
			// Buat reminder baru
			// Hitung sisa tagihan = Total - DP (asumsi sederhana, real case mungkin perlu tracking cicilan)
			// Disini kita ambil dari JumlahBayar original atau logic lain. 
			// Untuk simplifikasi: Kita buat reminder sebesar nominal bulanan (misal 1 juta)
			// Atau sisa tagihan. Kita pakai logic placeholder.
			
			// Ambil nominal bulanan dari Kamar via Booking
			var booking models.Pemesanan
			if err := s.db.Preload("Kamar").First(&booking, p.PemesananID).Error; err == nil {
				monthlyFee := booking.Kamar.HargaPerBulan
				
				reminder := models.PaymentReminder{
					PembayaranID:    p.ID,
					JumlahBayar:     monthlyFee,
					TanggalReminder: time.Now().AddDate(0, 0, 3), // H-3 dari jatuh tempo (logic disederhanakan)
					StatusReminder:  "Pending",
					IsSent:          false,
				}
				s.db.Create(&reminder)
				fmt.Printf("Created reminder for payment %d\n", p.ID)
			}
		}
	}
	
	return nil
}

// SendPendingReminders mengirim reminders yang sudah jatuh tempo
func (s *reminderService) SendPendingReminders() ([]models.PaymentReminder, error) {
	var reminders []models.PaymentReminder
	
	// Ambil reminder yang Pending, belum terkirim, dan tanggal reminder <= hari ini
	today := time.Now()
	if err := s.db.Where("status_reminder = ? AND is_sent = ? AND tanggal_reminder <= ?", "Pending", false, today).Find(&reminders).Error; err != nil {
		return nil, err
	}

	for i := range reminders {
		// Use Preloaded data to get tenant details
		var reminder models.PaymentReminder
		
		if err := s.db.Preload("Pembayaran.Pemesanan.Penyewa").Preload("Pembayaran.Pemesanan.Kamar").First(&reminder, reminders[i].ID).Error; err == nil {
			
			tenant := reminder.Pembayaran.Pemesanan.Penyewa
			kamar := reminder.Pembayaran.Pemesanan.Kamar
			
			// 1. Send Email
			if tenant.Email != "" {
				// Generate payment link (placeholder or real if available)
				paymentLink := fmt.Sprintf("http://localhost:3000/dashboard/payments/%d", reminder.PembayaranID)
				
				go func(email, name string, amount float64, due time.Time, link string) {
					s.emailSender.SendPaymentReminderEmail(email, name, amount, due, link)
				}(tenant.Email, tenant.NamaLengkap, reminder.JumlahBayar, reminder.Pembayaran.TanggalJatuhTempo, paymentLink)
			}
			
			// 2. Send WhatsApp
			if tenant.NomorHP != "" {
				msg := fmt.Sprintf("Halo %s, ini pengingat tagihan kos kamar %s sebesar Rp %.0f jatuh tempo pada %s. Mohon segera dibayar.", 
					tenant.NamaLengkap, kamar.NomorKamar, reminder.JumlahBayar, reminder.Pembayaran.TanggalJatuhTempo.Format("02 Jan 2006"))
				
				go func(phone, message string) {
					s.waSender.SendWhatsApp(phone, message)
				}(tenant.NomorHP, msg)
			}
			
			fmt.Printf("Sent notifications for Reminder ID %d\n", reminder.ID)
		}
		
		// Update status
		reminders[i].IsSent = true
		s.db.Save(&reminders[i])
	}
	
	return reminders, nil
}

// MarkReminderAsSent tandai reminder sudah dikirim
func (s *reminderService) MarkReminderAsSent(reminderID uint) error {
	return s.db.Model(&models.PaymentReminder{}).Where("id = ?", reminderID).Update("is_sent", true).Error
}

// MarkPaymentAsPaid tandai pembayaran reminder sebagai lunas
func (s *reminderService) MarkPaymentAsPaid(reminderID uint) error {
	return s.db.Model(&models.PaymentReminder{}).Where("id = ?", reminderID).Update("status_reminder", "Paid").Error
}

// GetPendingReminders ambil semua reminder yang pending
func (s *reminderService) GetPendingReminders() ([]models.PaymentReminder, error) {
	var reminders []models.PaymentReminder
	err := s.db.Where("status_reminder = ?", "Pending").Preload("Pembayaran.Pemesanan.Penyewa").Find(&reminders).Error
	return reminders, err
}
