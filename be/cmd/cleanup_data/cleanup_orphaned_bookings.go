package main

import (
	"fmt"
	"koskosan-be/internal/config"
	"koskosan-be/internal/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CleanupOrphanedBookings finds and fixes bookings with deleted rooms
// Run: go run cleanup_orphaned_bookings.go
func cleanupOrphanedBookings() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Connect to DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected. Starting orphaned bookings cleanup...")

	// 3. Find all bookings with deleted rooms
	var orphanedBookings []models.Pemesanan
	// Find bookings where the referenced kamar is soft-deleted OR kamar_id doesn't exist in kamars table
	if err := db.
		Where("kamar_id NOT IN (SELECT id FROM kamars WHERE deleted_at IS NULL)").
		Preload("Pembayaran").
		Preload("Penyewa").
		Find(&orphanedBookings).Error; err != nil {
		log.Fatalf("Failed to find orphaned bookings: %v", err)
	}

	if len(orphanedBookings) == 0 {
		log.Println("✅ No orphaned bookings found!")
		return
	}

	log.Printf("Found %d orphaned bookings. Processing...\n", len(orphanedBookings))

	// 4. Process each orphaned booking in a transaction
	for i, booking := range orphanedBookings {
		log.Printf("\n[%d/%d] Processing Booking ID: %d (Status: %s, Penyewa: %s)", 
			i+1, len(orphanedBookings), booking.ID, booking.StatusPemesanan, booking.Penyewa.NamaLengkap)

		err := db.Transaction(func(tx *gorm.DB) error {
			// A. Calculate booking status
			now := time.Now()
			isExpired := booking.TanggalKeluar.Before(now)
			hasPendingPayment := false
			totalPendingAmount := 0.0

			for _, p := range booking.Pembayaran {
				if p.StatusPembayaran == "Pending" && p.MetodePembayaran == "transfer" {
					hasPendingPayment = true
					totalPendingAmount += p.JumlahBayar
				}
			}

			// B. Handle based on booking status
			if isExpired {
				// Case 1: Booking sudah expired - soft delete saja
				log.Printf("  → Booking expired, soft deleting...")
				if err := tx.Model(&booking).Update("status_pemesanan", "Cancelled").Error; err != nil {
					return err
				}
				if err := tx.Model(&booking).Delete(&booking).Error; err != nil {
					return err
				}
				log.Printf("  ✓ Booking deleted (expired)")

			} else if booking.StatusPemesanan == "Pending" {
				// Case 2: Booking masih pending - batalkan
				log.Printf("  → Booking pending, cancelling...")
				if err := tx.Model(&booking).Update("status_pemesanan", "Cancelled").Error; err != nil {
					return err
				}
				log.Printf("  ✓ Booking cancelled")

				// Print info for manual follow-up if ada pending payment
				if hasPendingPayment {
					log.Printf("  ⚠️ WARNING: Booking had Rp %.0f pending transfer payment", totalPendingAmount)
					log.Printf("  Contact Info: %s | %s | %s", 
						booking.Penyewa.NomorHP, booking.Penyewa.Email, booking.Penyewa.NamaLengkap)
				}

			} else if booking.StatusPemesanan == "Confirmed" || booking.StatusPemesanan == "Aktif" || booking.StatusPemesanan == "Partially Paid" {
				// Case 3: Booking active - batalkan dan catat
				log.Printf("  → Booking active (%s), cancelling and logging for follow-up...", booking.StatusPemesanan)
				if err := tx.Model(&booking).Update("status_pemesanan", "Cancelled").Error; err != nil {
					return err
				}

				// Create audit log
				logEntry := fmt.Sprintf("ORPHANED_BOOKING_CLEANUP|BookingID:%d|OldStatus:%s|Reason:RoomDeleted|HasPendingPayment:%v|Amount:%.0f|Tenant:%s|Phone:%s",
					booking.ID, booking.StatusPemesanan, hasPendingPayment, totalPendingAmount, 
					booking.Penyewa.NamaLengkap, booking.Penyewa.NomorHP)
				log.Printf("  📝 Log: %s", logEntry)
				log.Printf("  ✓ Booking cancelled (active)")

				if hasPendingPayment {
					log.Printf("  🔴 CRITICAL: Booking had Rp %.0f pending payment! Follow-up needed!", totalPendingAmount)
					log.Printf("     Tenant: %s | Phone: %s | Email: %s", 
						booking.Penyewa.NamaLengkap, booking.Penyewa.NomorHP, booking.Penyewa.Email)
				}

			} else {
				// Case 4: Other status - just mark as cancelled
				log.Printf("  → Other status (%s), marking as cancelled...", booking.StatusPemesanan)
				if err := tx.Model(&booking).Update("status_pemesanan", "Cancelled").Error; err != nil {
					return err
				}
				log.Printf("  ✓ Booking status updated to Cancelled")
			}

			return nil
		})

		if err != nil {
			log.Printf("  ❌ Error processing booking %d: %v", booking.ID, err)
		}
	}

	log.Println("\n✅ Orphaned bookings cleanup completed!")
	log.Println("\n📋 SUMMARY:")
	log.Printf("   Total orphaned bookings processed: %d\n", len(orphanedBookings))
	log.Println("   Note: Check logs above for any critical items needing manual follow-up")
}
