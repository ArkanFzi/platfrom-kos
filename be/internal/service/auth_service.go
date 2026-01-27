package service

import (
	"errors"
	"koskosan-be/internal/config"
	"koskosan-be/internal/models"
	"koskosan-be/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, *models.User, error)
	Register(username, password, role string) (*models.User, error)
}

type authService struct {
	repo        repository.UserRepository
	penyewaRepo repository.PenyewaRepository
	config      *config.Config
}

func NewAuthService(repo repository.UserRepository, penyewaRepo repository.PenyewaRepository, cfg *config.Config) AuthService {
	return &authService{repo, penyewaRepo, cfg}
}

func (s *authService) Login(username, password string) (string, *models.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

func (s *authService) Register(username, password, role string) (*models.User, error) {
	// Check if user already exists
	_, err := s.repo.FindByUsername(username)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	// Create blank Penyewa record for tenants
	if role == "tenant" {
		penyewa := &models.Penyewa{
			UserID: user.ID,
		}
		if err := s.penyewaRepo.Create(penyewa); err != nil {
			// Note: We might want to handle this failure, maybe delete the user or just log it
			// For now we just log it or ignore
		}
	}

	return user, nil
}
