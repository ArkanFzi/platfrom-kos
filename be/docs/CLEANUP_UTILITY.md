# Data Cleanup Utility - Dokumentasi

## Overview

Utility untuk membersihkan data yang "stuck" atau "orphaned" ketika kamar dihapus tetapi masih ada booking yang referensi ke kamar tersebut.

## Problem

Ketika admin menghapus kamar:
1. Kamar dihapus dari database
2. Booking yang mereferensi kamar menjadi "orphaned" (tidak punya kamar valid)
3. User tidak bisa membatalkan booking karena data kamar sudah tidak ada
4. Data jadi "stuck" di sistem

## Solusi

Ada 2 opsi cleanup:

### 1. Cleanup Orphaned Bookings (RECOMMENDED) 🟢
```bash
cd /home/arkan/projects/rahmat_zaw/be
go run ./cmd/cleanup_data orphaned
```

**Apa yang dilakukan:**
- ✅ Menemukan semua booking yang kamarnya sudah dihapus
- ✅ Smart handling berdasarkan status booking:
  - **Booking expired** → Soft delete
  - **Booking pending** → Batalkan
  - **Booking active/confirmed** → Batalkan + catat untuk follow-up
- ✅ Deteksi pending transfer payments
- ✅ Log lengkap untuk audit trail

**Output Contoh:**
```
[1/5] Processing Booking ID: 123 (Status: Confirmed, Penyewa: Budi Santoso)
  → Booking active (Confirmed), cancelling and logging for follow-up...
  🔴 CRITICAL: Booking had Rp 500000 pending payment! Follow-up needed!
     Tenant: Budi Santoso | Phone: 081234567890 | Email: budi@email.com
  ✓ Booking cancelled (active)

✅ Orphaned bookings cleanup completed!
```

### 2. Full Reset (DESTRUCTIVE) 🔴
```bash
cd /home/arkan/projects/rahmat_zaw/be
go run ./cmd/cleanup_data all
```

**Apa yang dilakukan:**
- ⚠️ Hapus SEMUA Payment Reminders
- ⚠️ Hapus SEMUA Payments (Pembayaran)
- ⚠️ Hapus SEMUA Bookings (Pemesanan)
- ✅ Reset semua room status ke "Tersedia"

**⚠️ Gunakan hanya untuk:**
- Testing/development
- Full database reset
- BUKAN untuk production!

---

## Cara Menggunakan

### Step 1: Pastikan .env Sudah Konfigurasi
```bash
# Pastikan file /home/arkan/projects/rahmat_zaw/.env ada dan terisi dengan:
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=koskosan_rahmat
DB_PORT=5432
```

### Step 2: Run Cleanup
```bash
# Recommended: Clean orphaned bookings saja
cd /home/arkan/projects/rahmat_zaw/be
go run ./cmd/cleanup_data orphaned

# Atau jika perlu full reset
go run ./cmd/cleanup_data all
```

### Step 3: Check Output
Script akan menampilkan:
- ✓ Berapa banyak booking yang diproses
- ⚠️ Warnings untuk pending payments
- 🔴 Critical items yang perlu follow-up manual

---

## Skenario - Handling Berbeda

### Scenario 1: Booking Sudah Expired
```
Status: Expired
Action: Soft delete
Result: ✓ Booking deleted (expired)
```

### Scenario 2: Booking Pending (Belum Checkout)
```
Status: Pending, Active, Confirmed, Partially Paid
Action: Batalkan booking
Result: ✓ Booking cancelled
```

### Scenario 3: Booking Active + Ada Pending Payment
```
Status: Confirmed
Pending Payment: Rp 500,000 (Transfer, Pending)
Action: Batalkan + Log untuk follow-up
Result: 🔴 CRITICAL: Perlu kontak manual dengan penyewa
```

---

## Workflow Recommended

### Jika Ada Bookings Stuck:

1. **Jalankan cleanup:**
   ```bash
   go run ./cmd/cleanup_data orphaned
   ```

2. **Review output** untuk item yang critical:
   ```
   🔴 CRITICAL: Booking had Rp [amount] pending payment! Follow-up needed!
      Tenant: [Nama] | Phone: [No] | Email: [Email]
   ```

3. **Follow-up manual** untuk pending payments:
   - Hubungi tenant melalui nomor yang tercatat
   - Proses refund jika diperlukan
   - Update status pembayaran di sistem

4. **Verify** di database:
   ```sql
   -- Cek status booking setelah cleanup
   SELECT id, status_pemesanan, kamar_id FROM pemesanans WHERE status_pemesanan = 'Cancelled';
   ```

---

## Query untuk Detect Orphaned Bookings Manually

Jika ingin cek terlebih dahulu tanpa jalankan cleanup:

```sql
-- Find all bookings yang kamarnya sudah dihapus
SELECT 
    p.id as booking_id,
    p.status_pemesanan,
    p.kamar_id,
    pn.nama_lengkap as tenant_name,
    pn.nomor_hp,
    pn.email,
    p.created_at,
    COUNT(pb.id) as payment_count
FROM pemesanans p
LEFT JOIN penyewas pn ON p.penyewa_id = pn.id
LEFT JOIN pembayarans pb ON p.id = pb.pemesanan_id
WHERE p.kamar_id NOT IN (SELECT id FROM kamars WHERE deleted_at IS NULL)
GROUP BY p.id, pn.id
ORDER BY p.created_at DESC;
```

---

## Error Handling

| Error | Solusi |
|-------|--------|
| `Failed to connect to database` | Pastikan .env DB credentials benar |
| `database does not exist` | Pastikan database sudah dibuat |
| `relation "pemesanans" does not exist` | Pastikan migrations sudah jalan |

---

## Logs & Audit Trail

Semua cleanup dilog dengan detail:
- Booking ID
- Status perubahan
- Reason (RoomDeleted, Expired, etc)
- Pending payment info
- Tenant contact info

Gunakan untuk audit trail & follow-up manual jika diperlukan.

---

## Safety Notes

✅ **Orphaned cleanup aman** untuk production (hanya soft delete)
⚠️ **Full reset destructive** - backup database sebelum jalankan
✅ **Non-blocking** - tidak akan lock database
✅ **Transactional** - atomic operations, semua atau tidak sama sekali

---

## Support

Jika ada issue, check:
1. Database connectivity
2. Migrations status
3. .env configuration
4. Permission pada database
