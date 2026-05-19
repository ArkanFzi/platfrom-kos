#!/bin/bash

# Data Integrity Check Script
# Comprehensive checks for payment confirmation flow

set -e

if [ -f ".env" ]; then
    export $(cat .env | grep -v '#' | xargs)
fi

DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-3306}"
DB_USER="${DB_USER}"
DB_PASSWORD="${DB_PASSWORD}"
DB_NAME="${DB_NAME}"

echo "🔐 Data Integrity Check - Payment Confirmation Flow"
echo "=================================================="
echo ""

run_check() {
    local title="$1"
    local query="$2"
    
    echo "📋 $title"
    echo "   Query: $query"
    local result=$(mysql -h "${DB_HOST}" -u "${DB_USER}" -p"${DB_PASSWORD}" "${DB_NAME}" -N -e "$query" 2>&1)
    echo "   Result: $result"
    echo ""
}

# 1. Check database schema version
run_check "Database Schema Version (Columns in pembayaran table):" \
"SELECT GROUP_CONCAT(COLUMN_NAME) 
FROM INFORMATION_SCHEMA.COLUMNS 
WHERE TABLE_NAME='pembayaran' AND TABLE_SCHEMA='${DB_NAME}' 
ORDER BY ORDINAL_POSITION;"

# 2. Check for unique constraint on idempotency_key
run_check "Idempotency Key Constraint Status:" \
"SELECT CONSTRAINT_NAME, CONSTRAINT_TYPE 
FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS 
WHERE TABLE_NAME='pembayaran' AND TABLE_SCHEMA='${DB_NAME}' AND CONSTRAINT_NAME LIKE '%idempotency%';"

# 3. Check for existing indexes
run_check "Performance Indexes on pembayaran:" \
"SHOW INDEX FROM pembayaran WHERE Key_name NOT LIKE 'PRIMARY';"

# 4. Check foreign key relationships
run_check "Foreign Key Relationships:" \
"SELECT CONSTRAINT_NAME, COLUMN_NAME, REFERENCED_TABLE_NAME 
FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE 
WHERE TABLE_NAME='pembayaran' AND TABLE_SCHEMA='${DB_NAME}' AND REFERENCED_TABLE_NAME IS NOT NULL;"

# 5. Sample payment record to verify all fields
run_check "Sample Payment Record (newest 1):" \
"SELECT id, pemesanan_id, status_pembayaran, confirmed_at, idempotency_key, updated_at 
FROM pembayaran 
ORDER BY created_at DESC LIMIT 1;"

# 6. Check booking relationships
run_check "Booking-Payment Relationship Sample:" \
"SELECT p.id as payment_id, p.pemesanan_id, pe.status_pemesanan, p.status_pembayaran 
FROM pembayaran p 
LEFT JOIN pemesanan pe ON p.pemesanan_id = pe.id 
LIMIT 5;"

# 7. Check for cascade delete issues
run_check "Orphaned References Check:" \
"SELECT COUNT(*) as issue_count FROM pembayaran WHERE pemesanan_id NOT IN (SELECT id FROM pemesanan);"

# 8. Check confirmed_at vs status consistency
run_check "Confirmed Status vs confirmed_at Field Consistency:" \
"SELECT 
  COUNT(CASE WHEN status_pembayaran = 'Confirmed' AND confirmed_at IS NULL THEN 1 END) as missing_confirmation_time,
  COUNT(CASE WHEN status_pembayaran != 'Confirmed' AND confirmed_at IS NOT NULL THEN 1 END) as inconsistent_confirmation
FROM pembayaran;"

echo "✅ Integrity Check Complete!"
echo ""
echo "📊 Recommendations:"
echo "  1. If idempotency constraint is missing: Re-run migration 005"
echo "  2. If foreign keys show issues: Review cascade rules in database"
echo "  3. If consistency check shows problems: Run fix scripts"
echo "  4. If any count > 0: Contact DevOps team"
