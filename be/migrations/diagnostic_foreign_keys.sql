-- Diagnostic Query: Verify Foreign Key Constraints
-- Run this to ensure all relationships exist

-- 1. Check pembayaran → pemesanan relationship
SELECT 
    COUNT(*) as total_pembayaran,
    COUNT(DISTINCT pemesanan_id) as unique_bookings,
    COUNT(*) FILTER (WHERE pemesanan_id IS NULL) as null_booking_ids
FROM pembayaran;

-- 2. Check pemesanan → penyewa relationship
SELECT 
    COUNT(*) as total_pemesanan,
    COUNT(DISTINCT penyewa_id) as unique_tenants,
    COUNT(*) FILTER (WHERE penyewa_id IS NULL) as null_tenant_ids
FROM pemesanan;

-- 3. Check pemesanan → kamar relationship
SELECT 
    COUNT(*) as total_pemesanan,
    COUNT(DISTINCT kamar_id) as unique_rooms,
    COUNT(*) FILTER (WHERE kamar_id IS NULL) as null_room_ids
FROM pemesanan;

-- 4. Find orphaned payments
SELECT 'Orphaned Payments' as issue_type, COUNT(*) as count
FROM pembayaran p
WHERE NOT EXISTS (SELECT 1 FROM pemesanan WHERE id = p.pemesanan_id);

-- 5. Find orphaned bookings
SELECT 'Orphaned Bookings' as issue_type, COUNT(*) as count
FROM pemesanan p
WHERE NOT EXISTS (SELECT 1 FROM penyewa WHERE id = p.penyewa_id)
   OR NOT EXISTS (SELECT 1 FROM kamar WHERE id = p.kamar_id);

-- 6. Show confirmed payments without confirmed_at
SELECT 
    p.id, p.status_pembayaran, p.confirmed_at, p.updated_at,
    pe.status_pemesanan,
    k.nomor_kamar, k.status as kamar_status
FROM pembayaran p
LEFT JOIN pemesanan pe ON pe.id = p.pemesanan_id
LEFT JOIN kamar k ON k.id = pe.kamar_id
WHERE p.status_pembayaran = 'Confirmed' AND p.confirmed_at IS NULL
LIMIT 10;
