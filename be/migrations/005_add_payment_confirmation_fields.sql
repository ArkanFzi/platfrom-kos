-- Migration: Add payment confirmation tracking fields
-- Purpose: Track exact time of payment confirmation and prevent duplicate confirmations
-- Date: 2026-05-18
-- This migration adds fields that support FIX #1, #3, and #19 from payment service

-- Add confirmed_at column to track exact confirmation time for audit trail
ALTER TABLE pembayaran 
ADD COLUMN IF NOT EXISTS confirmed_at TIMESTAMP NULL;

-- Add comment to confirmed_at
COMMENT ON COLUMN pembayaran.confirmed_at IS 'Exact timestamp when payment was confirmed by admin (FIX #1, #3)';

-- Add idempotency_key column to prevent duplicate payment confirmations
ALTER TABLE pembayaran 
ADD COLUMN IF NOT EXISTS idempotency_key VARCHAR(255) UNIQUE NULL;

-- Add comment to idempotency_key
COMMENT ON COLUMN pembayaran.idempotency_key IS 'Unique key to prevent duplicate payment confirmations (FIX #19)';

-- Create index for idempotency_key lookups
CREATE INDEX IF NOT EXISTS idx_payment_idempotency ON pembayaran(idempotency_key);

-- Create index for confirmed_at queries (e.g., audit logs)
CREATE INDEX IF NOT EXISTS idx_payment_confirmed_at ON pembayaran(confirmed_at DESC);

-- Log migration completion
-- This migration is safe to run multiple times due to "IF NOT EXISTS" clauses
-- Expected execution time: < 500ms for tables with up to 100k records
