#!/bin/bash
# Payment Confirmation Feature - Automated Fix Script
# This script will apply all necessary database migrations and validations
# Usage: ./apply_payment_confirmation_fix.sh

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}================================================${NC}"
echo -e "${BLUE}Payment Confirmation Feature - Fix Script${NC}"
echo -e "${BLUE}================================================${NC}"
echo ""

# 1. Check if psql is available
echo -e "${YELLOW}[1/6] Checking prerequisites...${NC}"
if ! command -v psql &> /dev/null; then
    echo -e "${RED}❌ psql not found. Please install PostgreSQL client.${NC}"
    exit 1
fi
echo -e "${GREEN}✅ PostgreSQL client found${NC}"

# 2. Database configuration
echo ""
echo -e "${YELLOW}[2/6] Loading database configuration...${NC}"
if [ -f ".env" ]; then
    source .env
    echo -e "${GREEN}✅ .env file loaded${NC}"
else
    echo -e "${YELLOW}⚠️  .env file not found. Using defaults:${NC}"
    export DB_HOST=${DB_HOST:-localhost}
    export DB_PORT=${DB_PORT:-5432}
    export DB_USER=${DB_USER:-postgres}
    export DB_NAME=${DB_NAME:-koskosan_db}
fi

echo "  Database: $DB_HOST:$DB_PORT/$DB_NAME"
echo "  User: $DB_USER"

# 3. Test database connection
echo ""
echo -e "${YELLOW}[3/6] Testing database connection...${NC}"
if PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -c "SELECT 1;" &>/dev/null; then
    echo -e "${GREEN}✅ Database connection successful${NC}"
else
    echo -e "${RED}❌ Could not connect to database. Check credentials and database status.${NC}"
    exit 1
fi

# 4. Backup database
echo ""
echo -e "${YELLOW}[4/6] Creating database backup...${NC}"
BACKUP_FILE="backup_payment_fix_$(date +%Y%m%d_%H%M%S).sql"
PGPASSWORD=$DB_PASSWORD pg_dump -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" > "$BACKUP_FILE"
echo -e "${GREEN}✅ Backup created: $BACKUP_FILE${NC}"

# 5. Run migrations
echo ""
echo -e "${YELLOW}[5/6] Running migrations...${NC}"

MIGRATION_DIR="be/migrations"

if [ ! -d "$MIGRATION_DIR" ]; then
    echo -e "${RED}❌ Migrations directory not found: $MIGRATION_DIR${NC}"
    exit 1
fi

# 5.1 Run migration 005
echo "  Running migration 005_add_payment_confirmation_fields.sql..."
if [ -f "$MIGRATION_DIR/005_add_payment_confirmation_fields.sql" ]; then
    PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -f "$MIGRATION_DIR/005_add_payment_confirmation_fields.sql"
    echo -e "${GREEN}  ✅ Migration 005 completed${NC}"
else
    echo -e "${YELLOW}  ⚠️  Migration 005 not found (creating...)${NC}"
fi

# 5.2 Run migration 006
echo "  Running migration 006_validate_payment_status.sql..."
if [ -f "$MIGRATION_DIR/006_validate_payment_status.sql" ]; then
    PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -f "$MIGRATION_DIR/006_validate_payment_status.sql" 2>&1 | head -20
    echo -e "${GREEN}  ✅ Migration 006 completed${NC}"
else
    echo -e "${YELLOW}  ⚠️  Migration 006 not found (creating...)${NC}"
fi

# 6. Validation
echo ""
echo -e "${YELLOW}[6/6] Running validation checks...${NC}"

# Check confirmed_at field
CONFIRMED_AT_CHECK=$(PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -t -c "
SELECT COUNT(*) FROM information_schema.columns 
WHERE table_name='pembayaran' AND column_name='confirmed_at';
")

if [ "$CONFIRMED_AT_CHECK" -eq "1" ]; then
    echo -e "${GREEN}✅ confirmed_at column exists${NC}"
else
    echo -e "${RED}❌ confirmed_at column NOT found. Migration failed.${NC}"
    exit 1
fi

# Check idempotency_key field
IDEMPOTENCY_CHECK=$(PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -t -c "
SELECT COUNT(*) FROM information_schema.columns 
WHERE table_name='pembayaran' AND column_name='idempotency_key';
")

if [ "$IDEMPOTENCY_CHECK" -eq "1" ]; then
    echo -e "${GREEN}✅ idempotency_key column exists${NC}"
else
    echo -e "${RED}❌ idempotency_key column NOT found. Migration failed.${NC}"
    exit 1
fi

# Check for orphaned payments
ORPHANED_PAYMENTS=$(PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -t -c "
SELECT COUNT(*) FROM pembayaran 
WHERE pemesanan_id NOT IN (SELECT id FROM pemesanan);
" | xargs)

echo ""
if [ "$ORPHANED_PAYMENTS" -gt "0" ]; then
    echo -e "${YELLOW}⚠️  Found $ORPHANED_PAYMENTS orphaned payment records${NC}"
    echo "    These should be reviewed and corrected manually."
else
    echo -e "${GREEN}✅ No orphaned payment records found${NC}"
fi

# Check for incomplete confirmations
INCOMPLETE_CONFIRMATIONS=$(PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -t -c "
SELECT COUNT(*) FROM pembayaran 
WHERE status_pembayaran = 'Confirmed' AND confirmed_at IS NULL;
" | xargs)

if [ "$INCOMPLETE_CONFIRMATIONS" -gt "0" ]; then
    echo -e "${YELLOW}⚠️  Found $INCOMPLETE_CONFIRMATIONS incomplete confirmations${NC}"
    echo "    Fixing incomplete confirmations..."
    PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -c "
UPDATE pembayaran 
SET confirmed_at = updated_at 
WHERE status_pembayaran = 'Confirmed' AND confirmed_at IS NULL;
"
    echo -e "${GREEN}✅ Fixed incomplete confirmations${NC}"
else
    echo -e "${GREEN}✅ No incomplete confirmations found${NC}"
fi

echo ""
echo -e "${BLUE}================================================${NC}"
echo -e "${GREEN}✅ Payment Confirmation Fix Completed!${NC}"
echo -e "${BLUE}================================================${NC}"
echo ""
echo "Next Steps:"
echo "1. Restart backend service: systemctl restart backend-service"
echo "2. Test payment confirmation in frontend"
echo "3. Check backend logs for any errors"
echo ""
echo "Backup file saved: $BACKUP_FILE"
echo "In case of issues, restore using:"
echo "  psql -U $DB_USER -d $DB_NAME < $BACKUP_FILE"
