package service

import (
	"koskosan-be/internal/models"
	"koskosan-be/internal/repository"
)

type BookingResponse struct {
	ID              uint                `json:"id"`
	Kamar           models.Kamar        `json:"kamar"`
	TanggalMulai    string              `json:"tanggal_mulai"`
	DurasiSewa      int                 `json:"durasi_sewa"`
	StatusPemesanan string              `json:"status_pemesanan"`
	TotalBayar      float64             `json:"total_bayar"`
	StatusBayar     string              `json:"status_bayar"`
	Payments        []models.Pembayaran `json:"payments"`
}

type BookingService interface {
	GetUserBookings(userID uint) ([]BookingResponse, error)
}

type bookingService struct {
	repo        repository.BookingRepository
	penyewaRepo repository.PenyewaRepository
}

func NewBookingService(repo repository.BookingRepository, penyewaRepo repository.PenyewaRepository) BookingService {
	return &bookingService{repo, penyewaRepo}
}

func (s *bookingService) GetUserBookings(userID uint) ([]BookingResponse, error) {
	penyewa, err := s.penyewaRepo.FindByUserID(userID)
	if err != nil {
		return []BookingResponse{}, nil // No penyewa record yet
	}

	bookings, err := s.repo.FindByPenyewaID(penyewa.ID)
	if err != nil {
		return nil, err
	}

	var response []BookingResponse
	for _, b := range bookings {
		payments, _ := s.repo.GetPaymentsByBookingID(b.ID)

		var totalPaid float64
		var lastStatus string
		for _, p := range payments {
			if p.StatusPembayaran == "Confirmed" {
				totalPaid += p.JumlahBayar
			}
			lastStatus = p.StatusPembayaran
		}

		response = append(response, BookingResponse{
			ID:              b.ID,
			Kamar:           b.Kamar,
			TanggalMulai:    b.TanggalMulai.Format("2006-01-02"),
			DurasiSewa:      b.DurasiSewa,
			StatusPemesanan: b.StatusPemesanan,
			TotalBayar:      totalPaid,
			StatusBayar:     lastStatus,
			Payments:        payments,
		})
	}

	if response == nil {
		response = []BookingResponse{}
	}

	return response, nil
}
