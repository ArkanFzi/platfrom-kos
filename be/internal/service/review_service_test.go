package service

import (
	"errors"
	"koskosan-be/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockReviewRepository implements repository.ReviewRepository interface
type MockReviewRepository struct {
	mock.Mock
}

func (m *MockReviewRepository) Create(review *models.Ulasan) error {
	args := m.Called(review)
	return args.Error(0)
}

func (m *MockReviewRepository) FindByRoomID(kamarID uint) ([]models.Ulasan, error) {
	args := m.Called(kamarID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Ulasan), args.Error(1)
}

func (m *MockReviewRepository) FindAll() ([]models.Ulasan, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Ulasan), args.Error(1)
}

func (m *MockReviewRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test CreateReview - Success
func TestReviewService_CreateReview_Success(t *testing.T) {
	mockReviewRepo := new(MockReviewRepository)
	mockPaymentRepo := new(MockPaymentRepository)

	service := NewReviewService(mockReviewRepo, mockPaymentRepo)

	// Mock the payment verification
	payment := &models.Pembayaran{
		ID:               1,
		PemesananID:      1,
		StatusPembayaran: "Confirmed",
		TanggalBayar:     time.Now().Add(-7 * 24 * time.Hour), // 7 days ago
	}

	review := &models.Ulasan{
		KamarID:  1,
		UserID:   1,
		Rating:   5,
		Komentar: "Excellent room!",
	}

	mockPaymentRepo.On("FindByBookingID", mock.AnythingOfType("uint")).Return(payment, nil)
	mockReviewRepo.On("Create", mock.MatchedBy(func(r *models.Ulasan) bool {
		return r.KamarID == 1 && r.UserID == 1 && r.Rating == 5
	})).Return(nil)

	err := service.CreateReview(review, 1) // bookingID = 1

	assert.NoError(t, err)
	mockReviewRepo.AssertExpectations(t)
	mockPaymentRepo.AssertExpectations(t)
}

// Test CreateReview - No Confirmed Payment
func TestReviewService_CreateReview_NoConfirmedPayment(t *testing.T) {
	mockReviewRepo := new(MockReviewRepository)
	mockPaymentRepo := new(MockPaymentRepository)

	service := NewReviewService(mockReviewRepo, mockPaymentRepo)

	// Payment not found (user never booked this room)
	mockPaymentRepo.On("FindByBookingID", uint(1)).Return(nil, errors.New("record not found"))

	review := &models.Ulasan{
		KamarID:  1,
		UserID:   1,
		Rating:   5,
		Komentar: "Great!",
	}

	err := service.CreateReview(review, 1)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "must have confirmed payment")
	mockPaymentRepo.AssertExpectations(t)
	mockReviewRepo.AssertNotCalled(t, "Create")
}

// Test CreateReview - Invalid Rating
func TestReviewService_CreateReview_InvalidRating(t *testing.T) {
	mockReviewRepo := new(MockReviewRepository)
	mockPaymentRepo := new(MockPaymentRepository)

	service := NewReviewService(mockReviewRepo, mockPaymentRepo)

	payment := &models.Pembayaran{
		ID:               1,
		StatusPembayaran: "Confirmed",
	}

	mockPaymentRepo.On("FindByBookingID", uint(1)).Return(payment, nil)

	// Rating out of range (should be 1-5)
	review := &models.Ulasan{
		KamarID:  1,
		UserID:   1,
		Rating:   6, // Invalid
		Komentar: "Test",
	}

	err := service.CreateReview(review, 1)

	// Depends on implementation - might be validated in service or handler
	// Adjust assertion based on actual implementation
	assert.Error(t, err)
}

// Test GetReviewsByRoom - Success
func TestReviewService_GetReviewsByRoom_Success(t *testing.T) {
	mockReviewRepo := new(MockReviewRepository)
	mockPaymentRepo := new(MockPaymentRepository)

	service := NewReviewService(mockReviewRepo, mockPaymentRepo)

	expectedReviews := []models.Ulasan{
		{
			ID:       1,
			KamarID:  1,
			UserID:   1,
			Rating:   5,
			Komentar: "Excellent!",
		},
		{
			ID:       2,
			KamarID:  1,
			UserID:   2,
			Rating:   4,
			Komentar: "Good room",
		},
	}

	mockReviewRepo.On("FindByRoomID", uint(1)).Return(expectedReviews, nil)

	reviews, err := service.GetReviewsByRoom(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(reviews))
	assert.Equal(t, expectedReviews, reviews)
	mockReviewRepo.AssertExpectations(t)
}

// Test GetReviewsByRoom - Empty Result
func TestReviewService_GetReviewsByRoom_EmptyResult(t *testing.T) {
	mockReviewRepo := new(MockReviewRepository)
	mockPaymentRepo := new(MockPaymentRepository)

	service := NewReviewService(mockReviewRepo, mockPaymentRepo)

	emptyReviews := []models.Ulasan{}
	mockReviewRepo.On("FindByRoomID", uint(999)).Return(emptyReviews, nil)

	reviews, err := service.GetReviewsByRoom(999)

	assert.NoError(t, err)
	assert.NotNil(t, reviews)
	assert.Equal(t, 0, len(reviews))
	mockReviewRepo.AssertExpectations(t)
}

// Test GetReviewsByRoom - Repository Error
func TestReviewService_GetReviewsByRoom_Error(t *testing.T) {
	mockReviewRepo := new(MockReviewRepository)
	mockPaymentRepo := new(MockPaymentRepository)

	service := NewReviewService(mockReviewRepo, mockPaymentRepo)

	mockReviewRepo.On("FindByRoomID", uint(1)).Return(nil, errors.New("database error"))

	reviews, err := service.GetReviewsByRoom(1)

	assert.Error(t, err)
	assert.Nil(t, reviews)
	assert.Equal(t, "database error", err.Error())
	mockReviewRepo.AssertExpectations(t)
}

// Test GetAllReviews - Success
func TestReviewService_GetAllReviews_Success(t *testing.T) {
	mockReviewRepo := new(MockReviewRepository)
	mockPaymentRepo := new(MockPaymentRepository)

	service := NewReviewService(mockReviewRepo, mockPaymentRepo)

	expectedReviews := []models.Ulasan{
		{ID: 1, KamarID: 1, Rating: 5},
		{ID: 2, KamarID: 2, Rating: 4},
		{ID: 3, KamarID: 1, Rating: 3},
	}

	mockReviewRepo.On("FindAll").Return(expectedReviews, nil)

	reviews, err := service.GetAllReviews()

	assert.NoError(t, err)
	assert.Equal(t, 3, len(reviews))
	mockReviewRepo.AssertExpectations(t)
}

// Test DeleteReview - Success
func TestReviewService_DeleteReview_Success(t *testing.T) {
	mockReviewRepo := new(MockReviewRepository)
	mockPaymentRepo := new(MockPaymentRepository)

	service := NewReviewService(mockReviewRepo, mockPaymentRepo)

	mockReviewRepo.On("Delete", uint(1)).Return(nil)

	err := service.DeleteReview(1)

	assert.NoError(t, err)
	mockReviewRepo.AssertExpectations(t)
}

// Test DeleteReview - Not Found
func TestReviewService_DeleteReview_NotFound(t *testing.T) {
	mockReviewRepo := new(MockReviewRepository)
	mockPaymentRepo := new(MockPaymentRepository)

	service := NewReviewService(mockReviewRepo, mockPaymentRepo)

	mockReviewRepo.On("Delete", uint(999)).Return(errors.New("record not found"))

	err := service.DeleteReview(999)

	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())
	mockReviewRepo.AssertExpectations(t)
}
