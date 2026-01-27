package handlers

import (
	"koskosan-be/internal/models"
	"koskosan-be/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	service service.ReviewService
}

func NewReviewHandler(s service.ReviewService) *ReviewHandler {
	return &ReviewHandler{s}
}

func (h *ReviewHandler) GetReviews(c *gin.Context) {
	kamarIDStr := c.Param("id")
	kamarID, err := strconv.ParseUint(kamarIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Kamar ID"})
		return
	}

	reviews, err := h.service.GetReviewsByKamarID(uint(kamarID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if reviews == nil {
		reviews = []models.Review{}
	}
	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) GetAllReviews(c *gin.Context) {
	reviews, err := h.service.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if reviews == nil {
		reviews = []models.Review{}
	}
	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract UserID from context (set by AuthMiddleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	// Convert userID from interface{} to uint
	// Depending on how jwt claims are parsed, it might be float64 or uint
	switch v := userID.(type) {
	case float64:
		review.UserID = uint(v)
	case uint:
		review.UserID = v
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type in context"})
		return
	}

	if err := h.service.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, review)
}
