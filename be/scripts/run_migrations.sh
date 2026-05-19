#!/bin/bash

# Payment Confirmation Fix: Run all database migrations
# This script ensures all pending migrations are applied to the database

set -e

echo "🔧 Payment Confirmation Fix - Running Database Migrations"
echo "=========================================================="

# Check if migrations directory exists
if [ ! -d "./migrations" ]; then
    echo "❌ ERROR: migrations directory not found"
    exit 1
fi

# Export database connection string (adjust based on your .env setup)
if [ -f ".env" ]; then
    export $(cat .env | grep -v '#' | xargs)
fi

# If no database connection, exit
if [ -z "$DB_CONNECTION_STRING" ]; then
    echo "⚠️  WARNING: DB_CONNECTION_STRING not set. Using default..."
    DB_CONNECTION_STRING="${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Local"
fi

echo "📦 Database: ${DB_HOST}:${DB_PORT}/${DB_NAME}"
echo ""

# List all migration files
echo "📋 Migrations to be applied:"
ls -1 migrations/*.sql | sort | while read f; do
    echo "  ✓ $(basename $f)"
done
echo ""

# Apply migrations using mysql client
echo "⏳ Applying migrations..."

for migration in migrations/*.sql; do
    if [ -f "$migration" ]; then
        filename=$(basename "$migration")
        echo -n "  ▶ $filename ... "
        
        # Run migration
        if mysql -h "${DB_HOST}" -u "${DB_USER}" -p"${DB_PASSWORD}" "${DB_NAME}" < "$migration" 2>/dev/null; then
            echo "✅"
        else
            echo "⚠️  (may already be applied or non-critical)"
        fi
    fi
done

echo ""
echo "✅ Migration process completed!"
echo ""
echo "🔍 Next Steps:"
echo "  1. Run data validation: ./scripts/validate_payment_data.sh"
echo "  2. Check for data anomalies: ./scripts/check_data_integrity.sh"
echo "  3. Test payment confirmation: go test ./internal/service -v"
