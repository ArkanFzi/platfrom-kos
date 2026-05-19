package service

import (
	"errors"
	"koskosan-be/internal/models"
	"koskosan-be/internal/repository"
)

type KamarService interface {
	GetAll() ([]models.Kamar, error)
	GetByID(id uint) (*models.Kamar, error)
	Create(kamar *models.Kamar) error
	Update(kamar *models.Kamar) error
	Delete(id uint) error
	AddImage(image *models.KamarImage) error
	DeleteImagesByKamarID(kamarID uint) error
	CanDeleteRoom(id uint) (bool, string, error) // NEW: Check if room can be deleted
}

type kamarService struct {
	repo                repository.KamarRepository
	bookingRepo         repository.BookingRepository // NEW: Add booking repo to check active bookings
	paymentRepo         repository.PaymentRepository // NEW: Add payment repo to check pending payments
	notificationService NotificationService          // NEW: Add notification service for admin alerts
}

func NewKamarService(repo repository.KamarRepository, bookingRepo repository.BookingRepository, paymentRepo repository.PaymentRepository, notificationService NotificationService) KamarService {
	return &kamarService{
		repo:                repo,
		bookingRepo:         bookingRepo,
		paymentRepo:         paymentRepo,
		notificationService: notificationService,
	}
}

func (s *kamarService) GetAll() ([]models.Kamar, error) {
	return s.repo.FindAll()
}

func (s *kamarService) GetByID(id uint) (*models.Kamar, error) {
	return s.repo.FindByID(id)
}

func (s *kamarService) Create(kamar *models.Kamar) error {
	return s.repo.Create(kamar)
}

func (s *kamarService) Update(kamar *models.Kamar) error {
	// FIX #8: Prevent Admin from setting Room Status to Tersedia if occupied
	if kamar.Status == "Tersedia" {
		activeBooking, err := s.bookingRepo.FindActiveBookingByKamarID(kamar.ID)
		if err == nil && activeBooking != nil {
			// FEATURE #Eviction: Admin forced room to available. We must auto-cancel tying booking
			s.bookingRepo.UpdateStatus(activeBooking.ID, "Cancelled")
		}
	}
	return s.repo.Update(kamar)
}

func (s *kamarService) Delete(id uint) error {
	// Check if room can be deleted
	canDelete, reason, err := s.CanDeleteRoom(id)
	if err != nil {
		return err
	}
	if !canDelete {
		return errors.New(reason)
	}

	// NEW: Before deleting, check for any pending transfer payments for this room
	// This is to notify admin if they accidentally delete a room with pending payments
	pendingPayments, err := s.paymentRepo.FindPendingTransferPaymentsByRoomID(id)
	if err == nil && len(pendingPayments) > 0 {
		// Room is being deleted with pending transfer payments
		// Send notification to admin for each booking with pending payments
		for _, payment := range pendingPayments {
			booking, _ := s.bookingRepo.FindByID(payment.PemesananID)
			if booking != nil {
				// Send notification to admin
				_ = s.notificationService.NotifyAdminRoomDeletionWithPendingPayment(id, booking, []models.Pembayaran{payment})

				// Cancel the booking to prevent further issues
				_ = s.bookingRepo.UpdateStatus(booking.ID, "Cancelled")
			}
		}
	}

	return s.repo.Delete(id)
}

func (s *kamarService) CanDeleteRoom(id uint) (bool, string, error) {
	// Check if room has an active booking
	activeBooking, err := s.bookingRepo.FindActiveBookingByKamarID(id)
	if err != nil {
		return false, "", err
	}
	
	if activeBooking != nil {
		return false, "Kamar tidak dapat dihapus karena masih ada pemesanan aktif. Silakan ubah status kamar atau hubungi penyewa", nil
	}
	
	return true, "", nil
}

func (s *kamarService) AddImage(image *models.KamarImage) error {
	return s.repo.AddImage(image)
}

func (s *kamarService) DeleteImagesByKamarID(kamarID uint) error {
	return s.repo.DeleteImagesByKamarID(kamarID)
}
