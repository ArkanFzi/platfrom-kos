package validators

import (
	"errors"
	"mime/multipart"
	"regexp"
	"strings"
	"time"
)

var (
	// ErrInvalidEmail indicates invalid email format
	ErrInvalidEmail = errors.New("invalid email format")
	// ErrInvalidPhone indicates invalid phone number
	ErrInvalidPhone = errors.New("invalid phone number format")
	// ErrInvalidDateRange indicates invalid date range
	ErrInvalidDateRange = errors.New("invalid date range")
	// ErrInvalidPrice indicates invalid price
	ErrInvalidPrice = errors.New("invalid price")
	// ErrFileTooBig indicates file size exceeds limit
	ErrFileTooBig = errors.New("file size exceeds limit")
	// ErrInvalidFileType indicates unsupported file type
	ErrInvalidFileType = errors.New("invalid file type")
)

// Email regex pattern
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Phone regex pattern (Indonesian format)
var phoneRegex = regexp.MustCompile(`^(\+62|62|0)[0-9]{9,12}$`)

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	if email == "" || !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

// ValidatePhone validates phone number format
func ValidatePhone(phone string) error {
	if phone == "" || !phoneRegex.MatchString(phone) {
		return ErrInvalidPhone
	}
	return nil
}

// ValidateDateRange validates that start date is before end date
// and both are not in the past
func ValidateDateRange(start, end time.Time) error {
	now := time.Now()
	
	// Check if dates are in the past
	if start.Before(now.Add(-24 * time.Hour)) {
		return errors.New("start date cannot be in the past")
	}
	
	// Check if end is after start
	if !end.After(start) {
		return ErrInvalidDateRange
	}
	
	// Check reasonable max duration (e.g., 2 years)
	maxDuration := 730 * 24 * time.Hour // 2 years
	if end.Sub(start) > maxDuration {
		return errors.New("booking duration exceeds maximum allowed (2 years)")
	}
	
	return nil
}

// ValidatePrice validates price is positive and reasonable
func ValidatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}
	
	// Reasonable max: 100 million IDR per month
	if price > 100000000 {
		return errors.New("price exceeds reasonable maximum")
	}
	
	return nil
}

// ValidateFile validates file size and type
func ValidateFile(file *multipart.FileHeader, maxSizeMB int, allowedTypes []string) error {
	// Check file size
	maxBytes := int64(maxSizeMB * 1024 * 1024)
	if file.Size > maxBytes {
		return ErrFileTooBig
	}
	
	// Check file type
	fileType := strings.ToLower(file.Header.Get("Content-Type"))
	if len(allowedTypes) > 0 {
		allowed := false
		for _, t := range allowedTypes {
			if strings.Contains(fileType, t) {
				allowed = true
				break
			}
		}
		if !allowed {
			return ErrInvalidFileType
		}
	}
	
	return nil
}

// Common allowed types
var (
	// ImageTypes are allowed image MIME types
	ImageTypes = []string{"image/jpeg", "image/jpg", "image/png", "image/webp"}
	// DocumentTypes are allowed document MIME types
	DocumentTypes = []string{"application/pdf", "image/jpeg", "image/jpg", "image/png"}
)

// ValidateImageFile validates image file (common use case)
func ValidateImageFile(file *multipart.FileHeader) error {
	return ValidateFile(file, 5, ImageTypes) // 5MB max
}

// ValidatePaymentProof validates payment proof file
func ValidatePaymentProof(file *multipart.FileHeader) error {
	return ValidateFile(file, 5, DocumentTypes) // 5MB max
}

// ValidateUsername validates username format
func ValidateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if len(username) > 50 {
		return errors.New("username must not exceed 50 characters")
	}
	// Only alphanumeric and underscore
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username)
	if !matched {
		return errors.New("username can only contain letters, numbers, and underscores")
	}
	return nil
}

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	if len(password) > 100 {
		return errors.New("password must not exceed 100 characters")
	}
	return nil
}

// ValidateRequired validates that a string is not empty
func ValidateRequired(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New(fieldName + " is required")
	}
	return nil
}

// ValidatePositiveInt validates that an integer is positive
func ValidatePositiveInt(value int, fieldName string) error {
	if value <= 0 {
		return errors.New(fieldName + " must be positive")
	}
	return nil
}
