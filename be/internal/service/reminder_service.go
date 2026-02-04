package service

import (
	"fmt"
	"koskosan-be/internal/models"
	"koskosan-be/internal/repository"
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
	// Bisa ditambahkan service untuk email/notification di masa depan
}

func NewReminderService(paymentRepo repository.PaymentRepository) ReminderService {
	return &reminderService{paymentRepo}
}

// CreateMonthlyReminders membuat reminder untuk semua pembayaran DP yang jatuh tempo
func (s *reminderService) CreateMonthlyReminders() error {
	// Ambil semua pembayaran dengan tipe "dp" yang belum lunas
	// Query: SELECT * FROM pembayaran WHERE tipe_pembayaran = 'dp' AND status_pembayaran != 'Settled'
	
	// Untuk setiap pembayaran DP:
	// 1. Hitung jumlah yang harus dibayar = Total - DP
	// 2. Set tanggal reminder = 1 bulan setelah tanggal mulai sewa
	// 3. Buat payment reminder baru
	
	fmt.Println("Creating monthly reminders for DP payments")
	return nil
}

// SendPendingReminders mengirim reminders yang sudah jatuh tempo
func (s *reminderService) SendPendingReminders() ([]models.PaymentReminder, error) {
	// Ambil semua reminder yang:
	// 1. Status = "Pending"
	// 2. TanggalReminder <= hari ini
	// 3. IsSent = false
	
	// Untuk setiap reminder:
	// 1. Kirim notifikasi (email/SMS/push notification)
	// 2. Set IsSent = true
	
	fmt.Println("Sending pending reminders")
	return []models.PaymentReminder{}, nil
}

// MarkReminderAsSent tandai reminder sudah dikirim
func (s *reminderService) MarkReminderAsSent(reminderID uint) error {
	// Update reminder dengan ID = reminderID, set IsSent = true
	return nil
}

// MarkPaymentAsPaid tandai pembayaran reminder sebagai lunas
func (s *reminderService) MarkPaymentAsPaid(reminderID uint) error {
	// Update reminder dengan ID = reminderID, set StatusReminder = 'Paid'
	return nil
}

// GetPendingReminders ambil semua reminder yang pending
func (s *reminderService) GetPendingReminders() ([]models.PaymentReminder, error) {
	// Ambil semua reminder dengan status = 'Pending'
	return []models.PaymentReminder{}, nil
}

// Contoh penggunaan: Cron job yang dipanggil setiap hari untuk:
// 1. CreateMonthlyReminders() - untuk membuat reminder baru pada awal bulan
// 2. SendPendingReminders() - untuk mengirim reminder yang sudah jatuh tempo

// Scheduling example untuk digunakan di main.go dengan library seperti "github.com/robfig/cron"
// Bisa tambahkan:
/*
func ScheduleReminders(reminderService ReminderService) {
	c := cron.New()
	
	// Jalankan CreateMonthlyReminders setiap hari jam 00:00
	c.AddFunc("0 0 * * *", func() {
		reminderService.CreateMonthlyReminders()
	})
	
	// Jalankan SendPendingReminders setiap hari jam 08:00
	c.AddFunc("0 8 * * *", func() {
		reminderService.SendPendingReminders()
	})
	
	c.Start()
}
*/
