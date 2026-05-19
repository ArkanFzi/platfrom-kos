package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Pembayaran model (same as in models)
type Pembayaran struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	PemesananID      uint      `gorm:"index" json:"pemesanan_id"`
	JumlahBayar      float64   `json:"jumlah_bayar"`
	TanggalBayar     time.Time `json:"tanggal_bayar"`
	StatusPembayaran string    `gorm:"index" json:"status_pembayaran"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type PaymentStats struct {
	TotalPayments     int64   `json:"total_payments"`
	ConfirmedCount    int64   `json:"confirmed_count"`
	PendingCount      int64   `json:"pending_count"`
	RejectedCount     int64   `json:"rejected_count"`
	ConfirmedTotal    float64 `json:"confirmed_total"`
	PendingTotal      float64 `json:"pending_total"`
	APIResponseFormat string  `json:"api_response_format"`
	SamplePayments    []map[string]interface{} `json:"sample_payments"`
}

func main() {
	// Database connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "root"),
		getEnv("DB_NAME", "koskosan_rahmat"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Query statistics
	var totalPayments int64
	var confirmedCount int64
	var pendingCount int64
	var rejectedCount int64
	var confirmedTotal float64
	var pendingTotal float64

	// Total records
	db.Model(&Pembayaran{}).Count(&totalPayments)

	// Count by status
	db.Model(&Pembayaran{}).Where("status_pembayaran = ?", "Confirmed").Count(&confirmedCount)
	db.Model(&Pembayaran{}).Where("status_pembayaran = ?", "Pending").Count(&pendingCount)
	
	// Count rejected (can be either "Failed", "Rejected", or "Cancelled")
	var failedCount, rejectedStatusCount, cancelledCount int64
	db.Model(&Pembayaran{}).Where("status_pembayaran = ?", "Failed").Count(&failedCount)
	db.Model(&Pembayaran{}).Where("status_pembayaran = ?", "Rejected").Count(&rejectedStatusCount)
	db.Model(&Pembayaran{}).Where("status_pembayaran = ?", "Cancelled").Count(&cancelledCount)
	rejectedCount = failedCount + rejectedStatusCount + cancelledCount

	// Sum by status
	db.Model(&Pembayaran{}).Where("status_pembayaran = ?", "Confirmed").Select("COALESCE(SUM(jumlah_bayar), 0)").Row().Scan(&confirmedTotal)
	db.Model(&Pembayaran{}).Where("status_pembayaran = ?", "Pending").Select("COALESCE(SUM(jumlah_bayar), 0)").Row().Scan(&pendingTotal)

	// Sample payments
	var payments []Pembayaran
	db.Limit(5).Order("created_at DESC").Find(&payments)

	samplePayments := make([]map[string]interface{}, len(payments))
	for i, p := range payments {
		samplePayments[i] = map[string]interface{}{
			"id":                p.ID,
			"jumlah_bayar":      p.JumlahBayar,
			"status_pembayaran": p.StatusPembayaran,
			"tanggal_bayar":     p.TanggalBayar.Format("2006-01-02 15:04:05"),
			"pemesanan_id":      p.PemesananID,
		}
	}

	// Create response
	stats := PaymentStats{
		TotalPayments:     totalPayments,
		ConfirmedCount:    confirmedCount,
		PendingCount:      pendingCount,
		RejectedCount:     rejectedCount,
		ConfirmedTotal:    confirmedTotal,
		PendingTotal:      pendingTotal,
		APIResponseFormat: "raw_array",
		SamplePayments:    samplePayments,
	}

	// Output JSON
	jsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Println(string(jsonData))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
