# PAYMENT CONFIRMATION FEATURE - COMPREHENSIVE FIX GUIDE

## 🔧 STEP-BY-STEP FIXING PROCEDURE

### STEP 1: Database Migration (PRIORITY: CRITICAL)

#### 1.1 Run Migration untuk Confirmed Fields
```bash
cd /home/arkan/projects/rahmat_zaw/be

# Backup database first
pg_dump -U postgres -d koskosan_db > backup_$(date +%Y%m%d_%H%M%S).sql

# Run migration for payment confirmation fields
psql -U postgres -d koskosan_db -f migrations/005_add_payment_confirmation_fields.sql

# Verify migration
psql -U postgres -d koskosan_db -c "\d pembayaran"
# Should show: confirmed_at, idempotency_key columns
```

#### 1.2 Validate Payment Status
```bash
# Check for invalid status values
psql -U postgres -d koskosan_db -f migrations/diagnostic_foreign_keys.sql

# If found invalid data, they must be reviewed and corrected manually
# Contact admin for data review
```

### STEP 2: Backend Code Enhancement

#### 2.1 Update Payment Handler dengan Enhanced Logging
Replace [be/internal/handlers/payment_handler.go](be/internal/handlers/payment_handler.go) ConfirmPayment method:

```go
func (h *PaymentHandler) ConfirmPayment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.GlobalLogger.Error("Invalid payment ID: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	paymentID := uint(id)
	utils.GlobalLogger.Info("Confirming payment ID: %d", paymentID)

	if err := h.service.ConfirmPayment(paymentID); err != nil {
		utils.GlobalLogger.Error("Payment confirmation failed [ID:%d]: %v", paymentID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"payment_id": paymentID,
		})
		return
	}

	utils.GlobalLogger.Info("Payment confirmed successfully [ID:%d]", paymentID)
	c.JSON(http.StatusOK, gin.H{"message": "payment confirmed successfully"})
}
```

#### 2.2 Verify Payment Service Transaction Safety
Check [be/internal/service/payment_service.go](be/internal/service/payment_service.go#L47-L130):
- ✅ Uses `db.Transaction()` untuk atomicity
- ✅ Has idempotency check `if payment.StatusPembayaran == "Confirmed"`
- ✅ Implements pessimistic lock with `clause.Locking{Strength: "UPDATE"}`
- ✅ Updates all related records: payment, booking, kamar, penyewa, reminder

### STEP 3: Frontend Verification

#### 3.1 Test Payment Confirmation Component
Location: [fe/app/components/admin/PaymentConfirmation.tsx](fe/app/components/admin/PaymentConfirmation.tsx)

Check:
- ✅ `fetchPayments()` calls `api.getAllPayments()`
- ✅ `handleConfirm()` calls `api.confirmPayment(id)`
- ✅ Error toast shown on failure
- ✅ Payment list refreshed after confirmation

#### 3.2 Verify API Client
Location: [fe/app/services/api.ts](fe/app/services/api.ts#L467-L470)

```typescript
confirmPayment: async (paymentId: string) => {
    return apiCall<MessageResponse>('PUT', `/payments/${paymentId}/confirm`);
},
```

This should work correctly if backend responds with `{"message": "..."}`

### STEP 4: Manual Testing

#### 4.1 Create Test Payment
```bash
# 1. Login as admin
# 2. Go to Payment Management page
# 3. Find payment with status "Pending"
# 4. Click confirm button
# 5. Check browser console for any errors
# 6. Verify payment status changes to "Confirmed"
```

#### 4.2 Database Verification
```sql
-- After clicking confirm, check:
SELECT id, status_pembayaran, confirmed_at 
FROM pembayaran 
WHERE id = <payment_id>;

-- Should show:
-- status_pembayaran: 'Confirmed'
-- confirmed_at: current timestamp (not NULL)

-- Check related booking status changed
SELECT b.id, b.status_pemesanan, k.status as kamar_status
FROM pemesanan b
JOIN kamar k ON k.id = b.kamar_id
WHERE b.id = (SELECT pemesanan_id FROM pembayaran WHERE id = <payment_id>);

-- Should show:
-- status_pemesanan: 'Confirmed' (or 'Partially Paid' if DP)
-- kamar_status: 'Penuh'
```

### STEP 5: Common Issues & Solutions

| Issue | Symptom | Solution |
|-------|---------|----------|
| **Missing confirmed_at column** | `Error: column "confirmed_at" does not exist` | Run migration 005 |
| **Duplicate idempotency key** | `Error: duplicate key value violates unique constraint` | Run migration validation |
| **Orphaned payment records** | `Error: violates foreign key constraint` | Check diagnostic_foreign_keys.sql results |
| **Payment status not updating** | Status stays "Pending" after confirmation | Check transaction wasn't rolled back |
| **Kamar status not changing to Penuh** | Kamar still shows as "Booked" | Check pessimistic lock is acquired |
| **Penyewa role not updated to tenant** | Role stays "guest" after first payment | Confirm query counting payments correctly |

### STEP 6: Monitoring & Logs

#### 6.1 Check Backend Logs
```bash
# If running with docker
docker logs <backend-container> | grep -i "payment\|confirm" | tail -100

# If running locally
# Check console output for "Confirming payment ID" messages
```

#### 6.2 Check Database Logs (PostgreSQL)
```sql
-- Enable PostgreSQL query logging (if needed)
ALTER SYSTEM SET log_statement = 'all';
SELECT pg_reload_conf();

-- View logs in pg_log or journalctl
```

## 📋 VERIFICATION CHECKLIST

- [ ] Migration 005 executed successfully
- [ ] Database validation shows no orphaned records
- [ ] Payment handler logs showing confirmation attempts
- [ ] Test payment confirmation works end-to-end
- [ ] Database shows confirmed_at timestamp set
- [ ] Kamar status updated to "Penuh"
- [ ] Booking status updated to "Confirmed" or "Partially Paid"
- [ ] Penyewa role updated to "tenant" (first payment only)
- [ ] Email notification sent (if configured)
- [ ] WhatsApp notification sent (if configured)

## 🚨 EMERGENCY ROLLBACK

If issues arise:
```bash
# 1. Restore from backup
psql -U postgres -d koskosan_db < backup_20260518_HHMMSS.sql

# 2. Restart backend service
systemctl restart <backend-service-name>

# 3. Restart frontend
cd fe && npm run build && npm start
```

## 📞 CONTACT SUPPORT

If issues persist after following this guide:
1. Provide backend logs (last 100 lines)
2. Provide database query results from diagnostic_foreign_keys.sql
3. Provide exact error message from payment confirmation
4. Include payment ID that failed
