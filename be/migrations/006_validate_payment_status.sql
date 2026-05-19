-- Migration: Validate and normalize payment/booking status values
-- Purpose: Ensure database status values match enum definitions
-- Date: 2026-05-18

-- 1. PEMBAYARAN status validation and cleanup
-- Check for invalid pembayaran status values
SELECT 
    COUNT(*) as invalid_count,
    status_pembayaran
FROM pembayaran
WHERE status_pembayaran NOT IN ('Pending', 'Confirmed', 'Failed', 'Rejected', 'Settled', 'Cancelled')
GROUP BY status_pembayaran;

-- 2. PEMESANAN status validation and cleanup
-- Check for invalid pemesanan status values
SELECT 
    COUNT(*) as invalid_count,
    status_pemesanan
FROM pemesanan
WHERE status_pemesanan NOT IN ('Pending', 'Confirmed', 'Partially Paid', 'Active', 'Completed', 'Cancelled')
GROUP BY status_pemesanan;

-- 3. KAMAR status validation
-- Check for invalid kamar status values
SELECT 
    COUNT(*) as invalid_count,
    status
FROM kamar
WHERE status NOT IN ('Tersedia', 'Penuh', 'Booked', 'Perbaikan')
GROUP BY status;

-- 4. Check for orphaned payments (pembayaran tanpa pemesanan)
SELECT COUNT(*) as orphaned_payment_count
FROM pembayaran
WHERE pemesanan_id NOT IN (SELECT id FROM pemesanan);

-- 5. Check for payments with NULL confirmed_at for 'Confirmed' status
SELECT COUNT(*) as incomplete_confirmations
FROM pembayaran
WHERE status_pembayaran = 'Confirmed' AND confirmed_at IS NULL;

-- 6. Fix incomplete confirmations by setting confirmed_at to updated_at
UPDATE pembayaran
SET confirmed_at = updated_at
WHERE status_pembayaran = 'Confirmed' AND confirmed_at IS NULL;

-- NOTE: If queries above show invalid data:
-- Contact DevOps to run data cleanup scripts
-- Invalid statuses must be manually reviewed and corrected
