package service

import (
	"fmt"
	"koskosan-be/internal/config"
	"koskosan-be/internal/models"
	"koskosan-be/internal/repository"
	"koskosan-be/internal/utils"
)

type NotificationService interface {
	NotifyAdminRoomDeletionWithPendingPayment(kamarID uint, booking *models.Pemesanan, pendingPayments []models.Pembayaran) error
}

type notificationService struct {
	cfg         *config.Config
	waSender    utils.WhatsAppSender
	paymentRepo repository.PaymentRepository
	penyewaRepo repository.PenyewaRepository
}

func NewNotificationService(cfg *config.Config, waSender utils.WhatsAppSender, paymentRepo repository.PaymentRepository, penyewaRepo repository.PenyewaRepository) NotificationService {
	return &notificationService{
		cfg:         cfg,
		waSender:    waSender,
		paymentRepo: paymentRepo,
		penyewaRepo: penyewaRepo,
	}
}

// NotifyAdminRoomDeletionWithPendingPayment sends WhatsApp notification to admin when a room is deleted with pending transfer payments
// It includes user contact info, payment amount, and proof of transfer
func (s *notificationService) NotifyAdminRoomDeletionWithPendingPayment(kamarID uint, booking *models.Pemesanan, pendingPayments []models.Pembayaran) error {
	if s.cfg.AdminPhone == "" {
		return fmt.Errorf("ADMIN_PHONE is not configured")
	}

	if len(pendingPayments) == 0 {
		return nil // No pending payments to notify about
	}

	// Get penyewa (tenant) information
	penyewa, err := s.penyewaRepo.FindByID(booking.PenyewaID)
	if err != nil {
		return fmt.Errorf("failed to get penyewa info: %w", err)
	}

	// Build message with payment details
	message := fmt.Sprintf(
		"⚠️ NOTIFIKASI PENGHAPUSAN KAMAR DENGAN PEMBAYARAN PENDING\n\n"+
			"Admin telah menghapus kamar %s yang memiliki pemesanan dengan pembayaran yang menunggu konfirmasi.\n\n"+
			"📋 DATA PENYEWA:\n"+
			"Nama: %s\n"+
			"Nomor HP: %s\n"+
			"Email: %s\n\n",
		fmt.Sprintf("(ID: %d)", kamarID),
		penyewa.NamaLengkap,
		penyewa.NomorHP,
		penyewa.Email,
	)

	// Add payment details
	totalAmount := 0.0
	for i, payment := range pendingPayments {
		totalAmount += payment.JumlahBayar
		message += fmt.Sprintf(
			"💰 PEMBAYARAN #%d:\n"+
				"Jumlah: Rp %v\n"+
				"Status: %s\n"+
				"Metode: %s\n"+
				"Bukti Transfer: %s\n"+
				"Tanggal: %s\n\n",
			i+1,
			int64(payment.JumlahBayar),
			payment.StatusPembayaran,
			payment.MetodePembayaran,
			payment.BuktiTransfer,
			payment.CreatedAt.Format("02 January 2006 15:04"),
		)
	}

	message += fmt.Sprintf(
		"💵 TOTAL PEMBAYARAN YANG TERTUNDA: Rp %v\n\n"+
			"⚠️ TINDAKAN YANG DIPERLUKAN:\n"+
			"Silakan hubungi penyewa di nomor %s untuk mengkonfirmasi pengembalian dana atau menyelesaikan transaksi.\n"+
			"Verifikasi bukti transfer dan pastikan keamanan pengembalian uang kepada penyewa.\n\n"+
			"Tanggal Pelaporan: %s",
		int64(totalAmount),
		penyewa.NomorHP,
		booking.CreatedAt.Format("02 January 2006 15:04"),
	)

	// Send WhatsApp notification to admin
	if err := s.waSender.SendWhatsApp(s.cfg.AdminPhone, message); err != nil {
		return fmt.Errorf("failed to send whatsapp notification to admin: %w", err)
	}

	return nil
}
