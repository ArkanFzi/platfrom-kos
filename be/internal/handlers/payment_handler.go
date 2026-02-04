package handlers

import (
	"koskosan-be/internal/models"
	"koskosan-be/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(s service.PaymentService) *PaymentHandler {
	return &PaymentHandler{s}
}

func (h *PaymentHandler) GetAllPayments(c *gin.Context) {
	payments, err := h.service.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if payments == nil {
		payments = []models.Pembayaran{}
	}
	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) ConfirmPayment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	if err := h.service.ConfirmPayment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment confirmed successfully"})
}

func (h *PaymentHandler) CreateSnapToken(c *gin.Context) {
	var req struct {
		PemesananID   uint   `json:"pemesanan_id" binding:"required"`
		PaymentType   string `json:"payment_type" binding:"required"` // "full" atau "dp"
		PaymentMethod string `json:"payment_method" binding:"required"` // "midtrans" atau "cash"
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Validate payment type
	if req.PaymentType != "full" && req.PaymentType != "dp" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment type. Must be 'full' or 'dp'"})
		return
	}

	// Validate payment method
	if req.PaymentMethod != "midtrans" && req.PaymentMethod != "cash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment method. Must be 'midtrans' or 'cash'"})
		return
	}

	token, redirectURL, err := h.service.CreatePaymentSession(req.PemesananID, req.PaymentType, req.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":        token,
		"redirect_url": redirectURL,
		"payment_type": req.PaymentType,
		"payment_method": req.PaymentMethod,
	})
}

func (h *PaymentHandler) HandleMidtransWebhook(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	if err := h.service.HandleWebhook(payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook processed successfully"})
}
// ConfirmCashPayment untuk mengkonfirmasi pembayaran cash
func (h *PaymentHandler) ConfirmCashPayment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	var req struct {
		BuktiTransfer string `json:"bukti_transfer"` // Deskripsi/bukti transfer
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := h.service.ConfirmCashPayment(uint(id), req.BuktiTransfer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cash payment confirmed successfully"})
}