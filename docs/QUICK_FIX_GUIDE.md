# Payment Confirmation Fix - Quick Reference Guide

## 🚀 QUICK START (5 Minutes)

### Step 1: Apply Fixes
```bash
cd projects/rahmat_zaw/be
bash scripts/fix_payment_confirmation.sh
```

### Step 2: Verify Database
The script will automatically:
- ✅ Run migrations (005 & 006)
- ✅ Validate data integrity
- ✅ Fix incomplete confirmations
- ✅ Check schema consistency

### Step 3: Restart Backend
```bash
# If using Docker
docker-compose restart backend

# If running locally
go run ./cmd/main.go
```

### Step 4: Test in Admin Dashboard
1. Go to **Admin Dashboard** → **Payment Confirmation**
2. Click **Confirm** on a pending payment
3. Verify it shows **"Confirmed"** and update is successful

---

## 🔧 FIXES IMPLEMENTED

| FIX | Issue | Solution |
|-----|-------|----------|
| #1 | Room status race condition | Added pessimistic lock with SELECT...FOR UPDATE |
| #3 | Extend payment race condition | Added pessimistic lock in kamarRepo.FindByIDForUpdate() |
| #10 | DP payment not blocking move-in | Set booking status to "Partially Paid" instead of "Confirmed" |
| #11 | Extend payment not extending checkout | Calculate new checkout date: tanggal_mulai + (durasi_sewa + months) |
| #14 | Guests not promoted to tenants | Auto-promote on first confirmed payment |
| #19 | Duplicate payments possible | Added idempotency check: prevent confirm if already confirmed |

---

## 📋 FILES CREATED

### Migration Files
- `be/migrations/005_add_payment_confirmation_fields.sql` - Adds `confirmed_at` and `idempotency_key`
- `be/migrations/006_validate_payment_status.sql` - Validates enum values

### Script Files
- `be/scripts/run_migrations.sh` - Applies all database migrations
- `be/scripts/validate_payment_data.sh` - Validates and cleans data
- `be/scripts/check_data_integrity.sh` - Comprehensive data checks
- `be/scripts/fix_payment_confirmation.sh` - Master fix script
- `be/scripts/test_payment_confirmation.sh` - Test suite

### Service Files
- `be/internal/service/payment_confirmation_test.go` - Unit test cases
- `be/internal/service/payment_confirmation_logger.go` - Detailed logging

---

## 🔍 DEBUGGING ISSUES

### Issue: "Payment confirmation fails with error"

**Step 1**: Check logs for detailed error message
```bash
docker logs rahmat_zaw-backend | grep "PAYMENT_CONFIRMATION"
```

**Step 2**: Run data validation
```bash
bash be/scripts/validate_payment_data.sh
```

**Step 3**: Check database directly
```bash
# Check if payment record exists and has correct relations
SELECT p.id, p.status_pembayaran, p.confirmed_at, pe.id
FROM pembayaran p
LEFT JOIN pemesanan pe ON p.pemesanan_id = pe.id
WHERE p.id = YOUR_PAYMENT_ID;
```

### Issue: "Duplicate confirmations still possible"

Verify idempotency_key unique constraint exists:
```bash
SHOW INDEX FROM pembayaran WHERE Key_name LIKE '%idempotency%';
```

If missing, re-run migration 005.

### Issue: "Extend payment not extending checkout date"

Check if extend payment type is set correctly:
```bash
SELECT id, tipe_pembayaran, tanggal_jatuh_tempo
FROM pembayaran
WHERE id = YOUR_PAYMENT_ID;
```

Must have `tipe_pembayaran = 'extend'`.

---

## ✅ VERIFICATION CHECKLIST

After applying fixes, verify:

- [ ] Database migrations applied successfully
- [ ] `confirmed_at` column exists in pembayaran table
- [ ] `idempotency_key` has unique constraint
- [ ] No invalid status values in database
- [ ] Payment confirmation in UI works without errors
- [ ] Confirming same payment twice shows error (idempotency)
- [ ] Booking status updates correctly after payment
- [ ] Room status changes from "Tersedia" to "Penuh"
- [ ] DP payments set booking to "Partially Paid"
- [ ] Extend payments extend checkout date correctly
- [ ] Guest role promoted to "tenant" after first payment

---

## 📞 SUPPORT

If issues persist:

1. **Check logs**: `docker logs rahmat_zaw-backend`
2. **Run diagnostics**: `bash be/scripts/check_data_integrity.sh`
3. **Review migration**: `cat be/migrations/005_add_payment_confirmation_fields.sql`
4. **Check service code**: `cat be/internal/service/payment_service.go` (ConfirmPayment method)

---

## 📚 ADDITIONAL RESOURCES

- Full analysis: `PAYMENT_CONFIRMATION_FIX_GUIDE.md`
- Service implementation: `be/internal/service/payment_service.go`
- Database schema: `be/migrations/`
- Frontend component: `fe/app/components/admin/PaymentConfirmation.tsx`

---

**Last Updated**: 2026-05-19
**Status**: ✅ Ready for Production
