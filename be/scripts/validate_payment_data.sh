#!/bin/bash

# Payment Data Validation & Cleanup Script
# This script validates and fixes invalid payment/booking/kamar status values

set -e

if [ -f ".env" ]; then
    export $(cat .env | grep -v '#' | xargs)
fi

DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-3306}"
DB_USER="${DB_USER}"
DB_PASSWORD="${DB_PASSWORD}"
DB_NAME="${DB_NAME}"

echo "🔍 Payment Data Validation & Cleanup"
echo "===================================="
echo ""

# Function to run SQL query and display results
run_query() {
    local title="$1"
    local query="$2"
    
    echo "📊 $title"
    mysql -h "${DB_HOST}" -u "${DB_USER}" -p"${DB_PASSWORD}" "${DB_NAME}" -N -e "$query"
    echo ""
}

# 1. Check for invalid pembayaran status values
run_query "Invalid PEMBAYARAN Status Values:" \
"SELECT DISTINCT status_pembayaran, COUNT(*) as count 
FROM pembayaran 
WHERE status_pembayaran NOT IN ('Pending', 'Confirmed', 'Failed', 'Rejected', 'Settled', 'Cancelled')
GROUP BY status_pembayaran;"

# 2. Check for invalid pemesanan status values
run_query "Invalid PEMESANAN Status Values:" \
"SELECT DISTINCT status_pemesanan, COUNT(*) as count 
FROM pemesanan 
WHERE status_pemesanan NOT IN ('Pending', 'Confirmed', 'Partially Paid', 'Active', 'Completed', 'Cancelled')
GROUP BY status_pemesanan;"

# 3. Check for invalid kamar status values
run_query "Invalid KAMAR Status Values:" \
"SELECT DISTINCT status, COUNT(*) as count 
FROM kamar 
WHERE status NOT IN ('Tersedia', 'Penuh', 'Booked', 'Perbaikan')
GROUP BY status;"

# 4. Check for orphaned payments
run_query "Orphaned Payments (tanpa pemesanan):" \
"SELECT COUNT(*) as orphaned_count 
FROM pembayaran 
WHERE pemesanan_id NOT IN (SELECT id FROM pemesanan);"

# 5. Check for payments with NULL confirmed_at for 'Confirmed' status
run_query "Incomplete Confirmations (Confirmed status but no confirmed_at):" \
"SELECT COUNT(*) as incomplete_count 
FROM pembayaran 
WHERE status_pembayaran = 'Confirmed' AND confirmed_at IS NULL;"

echo "🔧 Auto-Fixing Data Issues..."
echo ""

# Fix 1: Set confirmed_at for Confirmed payments that don't have it
echo "▶ Fixing incomplete confirmations..."
mysql -h "${DB_HOST}" -u "${DB_USER}" -p"${DB_PASSWORD}" "${DB_NAME}" << EOF
UPDATE pembayaran
SET confirmed_at = updated_at
WHERE status_pembayaran = 'Confirmed' AND confirmed_at IS NULL;
EOF
echo "   ✅ Fixed incomplete confirmations"

# Fix 2: Check missing relations
echo ""
echo "▶ Checking relationship integrity..."
run_query "Booking Status Distribution:" \
"SELECT DISTINCT status_pemesanan, COUNT(*) as count 
FROM pemesanan 
GROUP BY status_pemesanan;"

run_query "Payment Status Distribution:" \
"SELECT DISTINCT status_pembayaran, COUNT(*) as count 
FROM pembayaran 
GROUP BY status_pembayaran;"

run_query "Room Status Distribution:" \
"SELECT DISTINCT status, COUNT(*) as count 
FROM kamar 
GROUP BY status;"

echo ""
echo "✅ Data Validation Complete!"
echo ""
echo "📝 Summary:"
echo "  • All invalid status values have been identified"
echo "  • Incomplete confirmations have been fixed"
echo "  • Relationship integrity has been checked"
echo ""
echo "⚠️  If invalid values remain above, please manually review:"
echo "  1. Check PAYMENT_CONFIRMATION_FIX_GUIDE.md for context"
echo "  2. Contact development team to review invalid data"
echo "  3. Run cleanup queries from 006_validate_payment_status.sql"
