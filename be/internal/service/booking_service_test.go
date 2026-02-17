package service

import (
	"errors"
	"koskosan-be/internal/models"
	"koskosan-be/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockBookingRepository implements repository.BookingRepository
type MockBookingRepository struct {
	mock.Mock
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func (m *MockBookingRepository) FindByPenyewaID(penyewaID uint) ([]models.Pemesanan, error) {
	args := m.Called(penyewaID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	return args.Get(0).([]models.Pemesanan), args.Error(1)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


// FindByPenyewaIDWithPayments - Added for Phase 3 optimization
func (m *MockBookingRepository) FindByPenyewaIDWithPayments(penyewaID uint) ([]models.Pemesanan, error) {
	args := m.Called(penyewaID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	return args.Get(0).([]models.Pemesanan), args.Error(1)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


// FindByID - Required by BookingRepository interface
func (m *MockBookingRepository) FindByID(id uint) (*models.Pemesanan, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	return args.Get(0).(*models.Pemesanan), args.Error(1)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func (m *MockBookingRepository) GetPaymentsByBookingID(bookingID uint) ([]models.Pembayaran, error) {
	args := m.Called(bookingID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func (m *MockBookingRepository) Create(booking *models.Pemesanan) error {
	args := m.Called(booking)
	return args.Error(0)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func (m *MockBookingRepository) Update(booking *models.Pemesanan) error {
	args := m.Called(booking)
	return args.Error(0)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func (m *MockBookingRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func (m *MockBookingRepository) WithTx(tx *gorm.DB) repository.BookingRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	return args.Get(0).(repository.BookingRepository)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func (m *MockBookingRepository) FindExpiredPendingBookings(beforeTime time.Time) ([]models.Pemesanan, error) {
	args := m.Called(beforeTime)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	return args.Get(0).([]models.Pemesanan), args.Error(1)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


// =============================================================================
// GET USER BOOKINGS TESTS
// =============================================================================

func TestGetUserBookings_Success(t *testing.T) {
	// Arrange
	mockBookingRepo := new(MockBookingRepository)
	mockPenyewaRepo := new(MockPenyewaRepository)

	penyewaID := uint(1)
	penyewa := &models.Penyewa{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	penyewa.ID = penyewaID

	// Phase 3 optimization: Payments are preloaded in Pembayaran relation
	bookings := []models.Pemesanan{
		{
			StatusPemesanan: "Pending",
			DurasiSewa:      30,
			Pembayaran: []models.Pembayaran{
				{StatusPembayaran: "Pending"}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
			}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
		}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
		{
			StatusPemesanan: "Confirmed",
			DurasiSewa:      60,
			Pembayaran: []models.Pembayaran{
				{StatusPembayaran: "Confirmed", JumlahBayar: 1000000}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
			}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
		}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	bookings[0].ID = 1
	bookings[1].ID = 2

	mockPenyewaRepo.On("FindByUserID", uint(1)).Return(penyewa, nil)
	// Using optimized method with preloaded payments
	mockBookingRepo.On("FindByPenyewaIDWithPayments", penyewaID).Return(bookings, nil)

	bookingService := NewBookingService(mockBookingRepo, mockPenyewaRepo, new(MockKamarRepository), new(MockPaymentRepository))

	// Act
	result, err := bookingService.GetUserBookings(1)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Pending", result[0].StatusPemesanan)
	assert.Equal(t, "Confirmed", result[1].StatusPemesanan)
	mockBookingRepo.AssertExpectations(t)
	mockPenyewaRepo.AssertExpectations(t)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func TestGetUserBookings_EmptyList(t *testing.T) {
	// Arrange
	mockBookingRepo := new(MockBookingRepository)
	mockPenyewaRepo := new(MockPenyewaRepository)

	penyewa := &models.Penyewa{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	penyewa.ID = 1

	mockPenyewaRepo.On("FindByUserID", uint(1)).Return(penyewa, nil)
	mockBookingRepo.On("FindByPenyewaIDWithPayments", uint(1)).Return([]models.Pemesanan{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
, nil)

	bookingService := NewBookingService(mockBookingRepo, mockPenyewaRepo, new(MockKamarRepository), new(MockPaymentRepository))

	// Act
	result, err := bookingService.GetUserBookings(1)

	// Assert
	assert.NoError(t, err)
	assert.Empty(t, result)
	mockBookingRepo.AssertExpectations(t)
	mockPenyewaRepo.AssertExpectations(t)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


// =============================================================================
// CREATE BOOKING TESTS
// =============================================================================

func TestCreateBooking_Success(t *testing.T) {
	// Arrange
	mockBookingRepo := new(MockBookingRepository)
	mockPenyewaRepo := new(MockPenyewaRepository)

	penyewa := &models.Penyewa{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	penyewa.ID = 1

	// User has NO active bookings
	mockPenyewaRepo.On("FindByUserID", uint(1)).Return(penyewa, nil)
	mockBookingRepo.On("FindByPenyewaID", uint(1)).Return([]models.Pemesanan{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
, nil)
	mockBookingRepo.On("Create", mock.AnythingOfType("*models.Pemesanan")).Return(nil)

	bookingService := NewBookingService(mockBookingRepo, mockPenyewaRepo, new(MockKamarRepository), new(MockPaymentRepository))

	// Act - Using actual signature: (userID uint, kamarID uint, tanggalMulai string, durasiSewa int)
	booking, err := bookingService.CreateBooking(1, 5, "2024-03-01", 30)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, booking)
	assert.Equal(t, uint(5), booking.KamarID)
	assert.Equal(t, "Pending", booking.StatusPemesanan)
	assert.Equal(t, 30, booking.DurasiSewa)
	mockBookingRepo.AssertExpectations(t)
	mockPenyewaRepo.AssertExpectations(t)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func TestCreateBooking_UserAlreadyHasActiveBooking(t *testing.T) {
	// Arrange
	mockBookingRepo := new(MockBookingRepository)
	mockPenyewaRepo := new(MockPenyewaRepository)

	penyewa := &models.Penyewa{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	penyewa.ID = 1

	// User already has a  PENDING booking (active)
	existingBookings := []models.Pemesanan{
		{StatusPemesanan: "Pending"}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


	mockPenyewaRepo.On("FindByUserID", uint(1)).Return(penyewa, nil)
	mockBookingRepo.On("FindByPenyewaID", uint(1)).Return(existingBookings, nil)

	bookingService := NewBookingService(mockBookingRepo, mockPenyewaRepo, new(MockKamarRepository), new(MockPaymentRepository))

	// Act
	booking, err := bookingService.CreateBooking(1, 1, "2024-03-15", 60)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, booking)
	assert.Contains(t, err.Error(), "pesanan aktif")
	mockBookingRepo.AssertExpectations(t)
	mockPenyewaRepo.AssertExpectations(t)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func TestCreateBooking_PenyewaNotFound(t *testing.T) {
	// Arrange
	mockBookingRepo := new(MockBookingRepository)
	mockPenyewaRepo := new(MockPenyewaRepository)

	mockPenyewaRepo.On("FindByUserID", uint(999)).Return(nil, errors.New("penyewa not found"))

	bookingService := NewBookingService(mockBookingRepo, mockPenyewaRepo, new(MockKamarRepository), new(MockPaymentRepository))

	// Act
	booking, err := bookingService.CreateBooking(999, 1, "2024-03-01", 30)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, booking)
	mockPenyewaRepo.AssertExpectations(t)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


func TestCreateBooking_InvalidDateFormat(t *testing.T) {
	// Arrange
	mockBookingRepo := new(MockBookingRepository)
	mockPenyewaRepo := new(MockPenyewaRepository)

	penyewa := &models.Penyewa{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

	penyewa.ID = 1

	mockPenyewaRepo.On("FindByUserID", uint(1)).Return(penyewa, nil)
	mockBookingRepo.On("FindByPenyewaID", uint(1)).Return([]models.Pemesanan{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
, nil)

	bookingService := NewBookingService(mockBookingRepo, mockPenyewaRepo, new(MockKamarRepository), new(MockPaymentRepository))

	// Act - Invalid date format
	booking, err := bookingService.CreateBooking(1, 1, "invalid-date", 30)

	// Assert
	assert.Error(t, err) // Should fail on time.Parse
	assert.Nil(t, booking)
	mockBookingRepo.AssertExpectations(t)
	mockPenyewaRepo.AssertExpectations(t)
}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


// =============================================================================
// TABLE-DRIVEN TESTS FOR BOOKING STATUS
// =============================================================================

func TestCreateBooking_DifferentStatuses(t *testing.T) {
	futureDate := time.Now().AddDate(0, 1, 0) // Next month
	pastDate := time.Now().AddDate(0, -2, 0)  // 2 months ago

	tests := []struct {
		name                  string
		existingBooking       models.Pemesanan
		shouldAllowNewBooking bool
		errorContains         string
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
{
		{
			name:                  "no existing bookings",
			existingBooking:       models.Pemesanan{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
, // Zero value, status ""
			shouldAllowNewBooking: true,
		}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
		{
			name:                  "existing Pending booking (active)",
			existingBooking:       models.Pemesanan{StatusPemesanan: "Pending"}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
			shouldAllowNewBooking: false,
			errorContains:         "pesanan aktif",
		}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
		{
			name: "existing Confirmed booking (active lease)",
			existingBooking: models.Pemesanan{
				StatusPemesanan: "Confirmed",
				TanggalMulai:    futureDate,
				DurasiSewa:      1,
			}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
			shouldAllowNewBooking: false,
			errorContains:         "masa sewa",
		}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
		{
			name: "existing Confirmed booking (expired lease)",
			existingBooking: models.Pemesanan{
				StatusPemesanan: "Confirmed",
				TanggalMulai:    pastDate,
				DurasiSewa:      1,
			}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
			shouldAllowNewBooking: true,
		}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
		{
			name:                  "existing Completed booking",
			existingBooking:       models.Pemesanan{StatusPemesanan: "Completed"}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
			shouldAllowNewBooking: true,
		}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
		{
			name:                  "existing Canceled booking",
			existingBooking:       models.Pemesanan{StatusPemesanan: "Canceled"}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
			shouldAllowNewBooking: true,
		}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
,
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBookingRepo := new(MockBookingRepository)
			mockPenyewaRepo := new(MockPenyewaRepository)

			penyewa := &models.Penyewa{}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

			penyewa.ID = 1

			var existingBookings []models.Pemesanan
			if tt.existingBooking.StatusPemesanan != "" {
				existingBookings = []models.Pemesanan{tt.existingBooking}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

			}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


			mockPenyewaRepo.On("FindByUserID", uint(1)).Return(penyewa, nil)
			mockBookingRepo.On("FindByPenyewaID", uint(1)).Return(existingBookings, nil)

			if tt.shouldAllowNewBooking {
				mockBookingRepo.On("Create", mock.AnythingOfType("*models.Pemesanan")).Return(nil)
			}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


			bookingService := NewBookingService(mockBookingRepo, mockPenyewaRepo, new(MockKamarRepository), new(MockPaymentRepository))

			// Act
			booking, err := bookingService.CreateBooking(1, 1, "2024-04-01", 30)

			// Assert
			if tt.shouldAllowNewBooking {
				assert.NoError(t, err)
				assert.NotNil(t, booking)
			}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
 else {
				assert.Error(t, err)
				assert.Nil(t, booking)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

			}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}


			mockBookingRepo.AssertExpectations(t)
			mockPenyewaRepo.AssertExpectations(t)
		}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}
)
	}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

}

// MockKamarRepository for testing
type MockKamarRepository struct {
	mock.Mock
}

func (m *MockKamarRepository) Create(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) Update(kamar *models.Kamar) error {
	args := m.Called(kamar)
	return args.Error(0)
}

func (m *MockKamarRepository) FindByID(id uint) (*models.Kamar, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) FindAll() ([]models.Kamar, error) {
	args := m.Called()
	return args.Get(0).([]models.Kamar), args.Error(1)
}

func (m *MockKamarRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockKamarRepository) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockKamarRepository) WithTx(tx *gorm.DB) repository.KamarRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.KamarRepository)
}

// MockPaymentRepository for testing
type  MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindByID(id uint) (*models.Pembayaran, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *models.Pembayaran) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) FindAll() ([]models.Pembayaran, error) {
	args := m.Called()
	return args.Get(0).([]models.Pembayaran), args.Error(1)
}

func (m *MockPaymentRepository) DeleteByBookingID(bookingID uint) error {
	args := m.Called(bookingID)
	return args.Error(0)
}

func (m *MockPaymentRepository) CreateReminder(reminder *models.PaymentReminder) error {
	args := m.Called(reminder)
	return args.Error(0)
}

func (m *MockPaymentRepository) WithTx(tx *gorm.DB) repository.PaymentRepository {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(repository.PaymentRepository)
}

