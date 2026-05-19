package main

import (
	"fmt"
	"log"
	"os"
	"koskosan-be/internal/config"
	"koskosan-be/internal/models"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Check command line arguments
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "orphaned":
		cleanupOrphanedBookings()
	case "all":
		fmt.Println("⚠️  WARNING: This will delete ALL bookings, payments, and reminders!")
		fmt.Println("Use 'orphaned' instead to clean up only orphaned bookings.")
		cleanupAllData()
	case "help":
		printUsage()
	default:
		fmt.Printf("❌ Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`
╔════════════════════════════════════════════════════════════════╗
║           KOSKOSAN DATA CLEANUP UTILITY                        ║
╚════════════════════════════════════════════════════════════════╝

USAGE:
  go run ./cmd/cleanup_data orphaned    - Clean orphaned bookings only
  go run ./cmd/cleanup_data all         - Delete ALL data (use with caution!)
  go run ./cmd/cleanup_data help        - Show this help message

COMMANDS:
  orphaned  - Remove bookings whose rooms have been deleted
              Handles: pending bookings, active bookings, payments
              Smart cleanup that protects valid data
              
  all       - DESTRUCTIVE: Delete ALL bookings, payments, reminders
              Use only for testing/reset
              Resets all room statuses to 'Tersedia'

EXAMPLES:
  # Clean orphaned bookings (RECOMMENDED)
  $ go run ./cmd/cleanup_data orphaned
  
  # Full reset (careful!)
  $ go run ./cmd/cleanup_data all

ENVIRONMENT:
  Make sure .env is configured with valid database credentials
  The script will connect using DB_* variables from .env

NOTES:
  - Orphaned cleanup is safe for production
  - All cleanup is logged for audit trail
  - Check output for any critical warnings
`)
}

func cleanupAllData() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Connect to DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected. Starting full cleanup...")

	// 3. Delete Data
	// Transaction to ensure atomicity
	err = db.Transaction(func(tx *gorm.DB) error {
		// A. Delete Payment Reminders
		if err := tx.Exec("DELETE FROM payment_reminders").Error; err != nil {
			return err
		}
		log.Println("Deleted all Payment Reminders")

		// B. Delete Payments (Pembayaran)
		// User said "clear atau ke pending". Assuming ALL for now to be safe/clean.
		if err := tx.Exec("DELETE FROM pembayarans").Error; err != nil {
			return err
		}
		log.Println("Deleted all Payments")

		// C. Delete Bookings (Pemesanan)
		if err := tx.Exec("DELETE FROM pemesanans").Error; err != nil {
			return err
		}
		log.Println("Deleted all Bookings")

		// D. Reset Room Status to 'Tersedia'
		if err := tx.Model(&models.Kamar{}).Where("1=1").Update("status", "Tersedia").Error; err != nil {
			return err
		}
		log.Println("Reset all Rooms to 'Tersedia'")

		return nil
	})

	if err != nil {
		log.Fatalf("Cleanup failed: %v", err)
	}

	log.Println("✅ Full cleanup completed successfully!")
}
