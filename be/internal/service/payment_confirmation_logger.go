package service

import (
	"fmt"
	"koskosan-be/internal/models"
	"log"
)

// PaymentConfirmationLogger provides detailed logging for debugging payment confirmation issues
type PaymentConfirmationLogger struct {
	prefix string
}

func NewPaymentConfirmationLogger() *PaymentConfirmationLogger {
	return &PaymentConfirmationLogger{prefix: "[PAYMENT_CONFIRMATION]"}
}

// LogConfirmationStart logs when a payment confirmation request starts
func (l *PaymentConfirmationLogger) LogConfirmationStart(paymentID uint) {
	log.Printf("%s 🔵 START: Confirming payment ID %d", l.prefix, paymentID)
}

// LogFetchPayment logs when fetching payment from database
func (l *PaymentConfirmationLogger) LogFetchPayment(payment *models.Pembayaran) {
	log.Printf("%s 📦 FETCH: Payment ID %d | Status: %s | Booking ID: %d | Amount: %.2f",
		l.prefix, payment.ID, payment.StatusPembayaran, payment.PemesananID, payment.JumlahBayar)
}

// LogIdempotencyCheck logs idempotency verification
func (l *PaymentConfirmationLogger) LogIdempotencyCheck(payment *models.Pembayaran, isAlreadyConfirmed bool) {
	if isAlreadyConfirmed {
		log.Printf("%s ⚠️  IDEMPOTENCY: Payment already confirmed at %s (FIX #19)",
			l.prefix, payment.ConfirmedAt.Format("2006-01-02 15:04:05"))
	} else {
		log.Printf("%s ✓ IDEMPOTENCY: Payment not yet confirmed, proceeding (FIX #19)", l.prefix)
	}
}

// LogStatusUpdate logs when payment status is updated
func (l *PaymentConfirmationLogger) LogStatusUpdate(oldStatus, newStatus string, confirmedAt string) {
	log.Printf("%s 🔄 STATUS: %s → %s | Confirmed at: %s (FIX #1)",
		l.prefix, oldStatus, newStatus, confirmedAt)
}

// LogReminderUpdate logs when payment reminder is updated
func (l *PaymentConfirmationLogger) LogReminderUpdate(paymentID uint, newStatus string) {
	log.Printf("%s 📩 REMINDER: Payment %d reminder status set to %s", l.prefix, paymentID, newStatus)
}

// LogBookingFetch logs when fetching related booking
func (l *PaymentConfirmationLogger) LogBookingFetch(booking *models.Pemesanan) {
	if booking == nil {
		log.Printf("%s ⚠️  BOOKING: Could not fetch booking (nil)", l.prefix)
		return
	}
	log.Printf("%s 📦 BOOKING: ID %d | Status: %s | Penyewa ID: %d | Kamar ID: %d",
		l.prefix, booking.ID, booking.StatusPemesanan, booking.PenyewaID, booking.KamarID)
}

// LogPaymentTypeHandling logs payment type specific handling
func (l *PaymentConfirmationLogger) LogPaymentTypeHandling(paymentType string, action string) {
	switch paymentType {
	case "dp":
		log.Printf("%s 💳 DOWN PAYMENT (DP): %s (FIX #10)", l.prefix, action)
	case "extend":
		log.Printf("%s 📅 EXTEND PAYMENT: %s (FIX #11)", l.prefix, action)
	case "full":
		log.Printf("%s 💰 FULL PAYMENT: %s", l.prefix, action)
	default:
		log.Printf("%s ❓ UNKNOWN TYPE (%s): %s", l.prefix, paymentType, action)
	}
}

// LogExtendPaymentCalculation logs extend payment calculation details
func (l *PaymentConfirmationLogger) LogExtendPaymentCalculation(oldDuration, newDuration int, oldCheckout, newCheckout string) {
	log.Printf("%s 📊 EXTEND CALC: Duration %d → %d months | Checkout %s → %s",
		l.prefix, oldDuration, newDuration, oldCheckout, newCheckout)
}

// LogBookingStatusUpdate logs booking status changes
func (l *PaymentConfirmationLogger) LogBookingStatusUpdate(bookingID uint, oldStatus, newStatus string) {
	log.Printf("%s 🔄 BOOKING: %s → %s (ID: %d) (FIX #3 with pessimistic lock)",
		l.prefix, oldStatus, newStatus, bookingID)
}

// LogRoomStatusUpdate logs room status changes
func (l *PaymentConfirmationLogger) LogRoomStatusUpdate(roomID uint, oldStatus, newStatus string) {
	log.Printf("%s 🏠 ROOM: %s → %s (ID: %d) | Using pessimistic lock for atomicity (FIX #1)",
		l.prefix, oldStatus, newStatus, roomID)
}

// LogTenantPromotion logs when guest is promoted to tenant
func (l *PaymentConfirmationLogger) LogTenantPromotion(penyewaID uint) {
	log.Printf("%s 👤 PROMOTION: Penyewa %d promoted from guest → tenant (FIX #14)", l.prefix, penyewaID)
}

// LogNotificationStart logs when sending notifications
func (l *PaymentConfirmationLogger) LogNotificationStart(paymentID uint) {
	log.Printf("%s 📬 NOTIFY: Starting notification process for payment %d", l.prefix, paymentID)
}

// LogConfirmationSuccess logs when confirmation completes successfully
func (l *PaymentConfirmationLogger) LogConfirmationSuccess(paymentID uint) {
	log.Printf("%s ✅ SUCCESS: Payment %d confirmed successfully", l.prefix, paymentID)
}

// LogConfirmationError logs when confirmation fails
func (l *PaymentConfirmationLogger) LogConfirmationError(paymentID uint, err error) {
	log.Printf("%s ❌ ERROR: Payment %d confirmation failed: %v", l.prefix, paymentID, err)
}

// LogTransactionRollback logs when transaction is rolled back
func (l *PaymentConfirmationLogger) LogTransactionRollback(paymentID uint, reason error) {
	log.Printf("%s 🔙 ROLLBACK: Transaction rolled back for payment %d | Reason: %v", l.prefix, paymentID, reason)
}

// LogDatabaseQuery logs raw database queries for debugging (use cautiously in production)
func (l *PaymentConfirmationLogger) LogDatabaseQuery(queryName string, details string) {
	log.Printf("%s 🗄️  DB: %s - %s", l.prefix, queryName, details)
}

// LogWarning logs warnings that don't stop the process
func (l *PaymentConfirmationLogger) LogWarning(message string) {
	log.Printf("%s ⚠️  WARNING: %s", l.prefix, message)
}

// SummarizeConfirmationFlow logs a complete flow summary for debugging
func (l *PaymentConfirmationLogger) SummarizeConfirmationFlow(payment *models.Pembayaran, booking *models.Pemesanan, wasSuccessful bool, errorMsg string) {
	summary := fmt.Sprintf(`
%s 📋 CONFIRMATION FLOW SUMMARY:
  ├─ Payment ID: %d
  ├─ Status: %s → %s
  ├─ Booking ID: %d
  ├─ Booking Status: %s
  ├─ Amount: %.2f
  ├─ Payment Type: %s
  ├─ Confirmed At: %s
  ├─ Result: %v
  └─ Error: %s
`,
		l.prefix,
		payment.ID,
		"Pending",
		payment.StatusPembayaran,
		booking.ID,
		booking.StatusPemesanan,
		payment.JumlahBayar,
		payment.TipePembayaran,
		payment.ConfirmedAt.Format("2006-01-02 15:04:05"),
		wasSuccessful,
		errorMsg,
	)
	log.Println(summary)
}
