package handlers

import (
	"koskosan-be/internal/config"
	"koskosan-be/internal/service"
	"koskosan-be/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
	cfg     *config.Config
}

func NewAuthHandler(s service.AuthService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		service: s,
		cfg:     cfg,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Sanitize input
	input.Username = utils.SanitizeString(input.Username)
	input.Password = utils.SanitizeString(input.Password)

	token, user, err := h.service.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set token as httpOnly cookie with secure flag in production
	c.SetCookie("token", token, 3600, "/", "", h.cfg.IsProduction, true)

	c.JSON(http.StatusOK, gin.H{
		"token": token, // Also return token for frontend localStorage
		"user":  user,
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
		Birthdate string `json:"birthdate"` // "YYYY-MM-DD"
	}

	if err := c.ShouldBindJSON(&input); err  != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Sanitize and validate input
	input.Username = utils.SanitizeString(input.Username)
	input.Password = utils.SanitizeString(input.Password)
	input.Email = utils.SanitizeString(input.Email)
	input.Phone = utils.SanitizeString(input.Phone)
	input.Address = utils.SanitizeString(input.Address)
	input.Birthdate = utils.SanitizeString(input.Birthdate)

	// Validate username
	if validationErr := utils.ValidateUsername(input.Username); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Message})
		return
	}

	// Validate password strength
	if validationErr := utils.ValidatePassword(input.Password); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Message})
		return
	}

	// Parse birthdate
	
	user, err := h.service.Register(input.Username, input.Password, "tenant", input.Email, input.Phone, input.Address, input.Birthdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Picture  string `json:"picture"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Validate email format
	if validationErr := utils.ValidateEmail(input.Email); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Message})
		return
	}

	// Sanitize inputs
	input.Email = utils.SanitizeString(input.Email)
	input.Username = utils.SanitizeString(input.Username)
	input.Picture = utils.SanitizeString(input.Picture)

	token, user, err := h.service.GoogleLogin(input.Email, input.Username, input.Picture)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	err := h.service.ForgotPassword(input.Email)
	if err != nil {
		// In a real generic app, we shouldn't return detailed user-not-found errors to public.
		// But for tenant/friendly apps, it's often accepted.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link has been sent to your email"})
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var input struct {
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if validationErr := utils.ValidatePassword(input.NewPassword); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Message})
		return
	}

	if err := h.service.ResetPassword(input.Token, input.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password has been reset successfully"})
}
