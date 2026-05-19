package repository

import (
	"koskosan-be/internal/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	FindAll() ([]models.Pembayaran, error)
	FindByID(id uint) (*models.Pembayaran, error)
	FindByOrderID(orderID string) (*models.Pembayaran, error)
	FindByBookingID(bookingID uint) ([]models.Pembayaran, error)
	FindPendingTransferPaymentsByBookingID(bookingID uint) ([]models.Pembayaran, error) // NEW: Get pending transfer payments for a booking
	FindPendingTransferPaymentsByRoomID(kamarID uint) ([]models.Pembayaran, error)      // NEW: Get all pending transfer payments for a room
	Create(payment *models.Pembayaran) error
	CreateReminder(reminder *models.PaymentReminder) error
	Update(payment *models.Pembayaran) error
	DeleteByBookingID(bookingID uint) error
	DeleteRemindersByBookingID(bookingID uint) error
	WithTx(tx *gorm.DB) PaymentRepository
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) FindAll() ([]models.Pembayaran, error) {
	var payments []models.Pembayaran
	err := r.db.Preload("Pemesanan.Penyewa.User").Preload("Pemesanan.Kamar").Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	var payment models.Pembayaran
	err := r.db.Preload("Pemesanan.Penyewa.User").Preload("Pemesanan.Kamar").First(&payment, id).Error
	return &payment, err
}

func (r *paymentRepository) FindByOrderID(orderID string) (*models.Pembayaran, error) {
	var payment models.Pembayaran
	err := r.db.Preload("Pemesanan.Penyewa.User").Preload("Pemesanan.Kamar").Where("order_id = ?", orderID).First(&payment).Error
	return &payment, err
}

func (r *paymentRepository) FindByBookingID(bookingID uint) ([]models.Pembayaran, error) {
	var payments []models.Pembayaran
	err := r.db.Where("pemesanan_id = ?", bookingID).Order("created_at DESC").Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) FindPendingTransferPaymentsByBookingID(bookingID uint) ([]models.Pembayaran, error) {
	var payments []models.Pembayaran
	// Find payments with status "Pending" and metode_pembayaran "transfer"
	err := r.db.
		Preload("Pemesanan.Penyewa.User").
		Preload("Pemesanan.Kamar").
		Where("pemesanan_id = ? AND status_pembayaran = ? AND metode_pembayaran = ?", bookingID, "Pending", "transfer").
		Order("created_at DESC").
		Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) FindPendingTransferPaymentsByRoomID(kamarID uint) ([]models.Pembayaran, error) {
	var payments []models.Pembayaran
	// Find all pending transfer payments for a specific room
	// Join through pemesanan table to match kamar_id
	err := r.db.
		Joins("INNER JOIN pemesanans ON pembayarans.pemesanan_id = pemesanans.id").
		Preload("Pemesanan.Penyewa.User").
		Preload("Pemesanan.Kamar").
		Where("pemesanans.kamar_id = ? AND pembayarans.status_pembayaran = ? AND pembayarans.metode_pembayaran = ?", kamarID, "Pending", "transfer").
		Order("pembayarans.created_at DESC").
		Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) Create(payment *models.Pembayaran) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) Update(payment *models.Pembayaran) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	return r.db.Create(reminder).Error
}

func (r *paymentRepository) WithTx(tx *gorm.DB) PaymentRepository {
	return &paymentRepository{db: tx}
}

func (r *paymentRepository) DeleteByBookingID(bookingID uint) error {
	// First delete associated reminders, then delete payments
	if err := r.DeleteRemindersByBookingID(bookingID); err != nil {
		return err
	}
	return r.db.Where("pemesanan_id = ?", bookingID).Delete(&models.Pembayaran{}).Error
}

func (r *paymentRepository) DeleteRemindersByBookingID(bookingID uint) error {
	// Delete all reminders linked to payments of this booking
	return r.db.Where(
		"pembayaran_id IN (SELECT id FROM pembayarans WHERE pemesanan_id = ?)",
		bookingID,
	).Delete(&models.PaymentReminder{}).Error
}
