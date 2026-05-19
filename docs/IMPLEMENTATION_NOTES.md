# Fitur Proteksi Penghapusan Kamar dengan Pending Payment - Dokumentasi Implementasi

## Overview
Fitur ini melindungi penyewa dari kehilangan pembayaran mereka jika admin secara tidak sengaja menghapus kamar yang memiliki pembayaran transfer yang sedang menunggu konfirmasi admin.

## Alur Kerja

### Skenario:
1. Penyewa membuat pemesanan dengan metode pembayaran **transfer**
2. Penyewa melakukan transfer dan upload bukti transfer (status pembayaran: "Pending")
3. Admin belum mengkonfirmasi pembayaran saat itu
4. Admin secara tidak sengaja menghapus kamar tersebut
5. **Sistem akan:**
   - Mendeteksi adanya pembayaran transfer yang pending
   - Mengirim notifikasi WhatsApp ke admin dengan:
     - Nama lengkap dan nomor HP penyewa
     - Email penyewa
     - Jumlah uang yang telah ditransfer
     - Bukti transfer (URL/path file)
     - Tanggal pemesanan
   - Otomatis membatalkan pemesanan (status: "Cancelled")

### Fitur Ini HANYA Bekerja Untuk:
- Metode pembayaran: **TRANSFER** (bukan tunai/cash)
- Status pembayaran: **PENDING** (belum dikonfirmasi admin)

---

## Perubahan Kode

### 1. **Config (internal/config/config.go)**
Menambahkan field baru untuk menyimpan nomor WhatsApp admin:
```go
AdminPhone string // Admin WhatsApp number for notifications
```

File `.env.example` juga diperbarui untuk menambahkan:
```
ADMIN_PHONE=62812345678
```

### 2. **Payment Repository (internal/repository/payment_repository.go)**
Menambahkan 2 method baru:

#### `FindPendingTransferPaymentsByBookingID(bookingID uint)`
- Mencari pembayaran transfer dengan status "Pending" untuk booking tertentu
- Preload semua relasi yang diperlukan (penyewa, kamar, user)

#### `FindPendingTransferPaymentsByRoomID(kamarID uint)`
- Mencari semua pembayaran transfer dengan status "Pending" untuk kamar tertentu
- Menggunakan INNER JOIN dengan tabel pemesanan untuk mendapatkan kamar_id
- Preload semua relasi yang diperlukan

### 3. **Notification Service (internal/service/notification_service.go)** - FILE BARU
Service baru untuk menangani notifikasi WhatsApp ke admin:

#### Interface:
```go
type NotificationService interface {
    NotifyAdminRoomDeletionWithPendingPayment(kamarID uint, booking *models.Pemesanan, pendingPayments []models.Pembayaran) error
}
```

#### Fitur Utama:
- Mengumpulkan informasi penyewa dan pembayaran
- Format pesan WhatsApp yang informatif dan terstruktur
- Mengirim notifikasi melalui Fonnte WhatsApp API
- Menangani error dengan graceful

#### Struktur Pesan WhatsApp:
```
⚠️ NOTIFIKASI PENGHAPUSAN KAMAR DENGAN PEMBAYARAN PENDING

Admin telah menghapus kamar (ID: X) yang memiliki pemesanan dengan pembayaran yang menunggu konfirmasi.

📋 DATA PENYEWA:
Nama: [Nama Lengkap]
Nomor HP: [Nomor]
Email: [Email]

💰 PEMBAYARAN:
Jumlah: Rp [Jumlah]
Status: Pending
Metode: transfer
Bukti Transfer: [URL/Path]
Tanggal: [Tanggal Pemesanan]

💵 TOTAL PEMBAYARAN YANG TERTUNDA: Rp [Total]

⚠️ TINDAKAN YANG DIPERLUKAN:
Silakan hubungi penyewa di nomor [Nomor] untuk mengkonfirmasi pengembalian dana atau menyelesaikan transaksi.
```

### 4. **Kamar Service (internal/service/kamar_service.go)**
Memperbarui struct dan method:

#### Struct Update:
```go
type kamarService struct {
    repo                repository.KamarRepository
    bookingRepo         repository.BookingRepository
    paymentRepo         repository.PaymentRepository         // NEW
    notificationService NotificationService                  // NEW
}
```

#### Constructor Update:
```go
func NewKamarService(repo repository.KamarRepository, 
    bookingRepo repository.BookingRepository, 
    paymentRepo repository.PaymentRepository,             // NEW
    notificationService NotificationService) KamarService // NEW
```

#### Delete Method Logic:
1. Validasi apakah kamar bisa dihapus (via `CanDeleteRoom`)
2. Cari semua pembayaran transfer pending untuk kamar ini
3. Jika ada pembayaran pending:
   - Kirim notifikasi WhatsApp ke admin untuk setiap pembayaran
   - Otomatis batalkan pemesanan terkait (status: "Cancelled")
4. Lanjutkan proses penghapusan kamar
5. Jika ada error saat mencari payment atau mengirim notifikasi, tetap lanjutkan deletion (non-blocking)

### 5. **Main Entry Point (cmd/api/main.go)**
Menambahkan inisialisasi NotificationService:

```go
// Inisialisasi NotificationService sebelum KamarService
notificationService := service.NewNotificationService(cfg, waSender, paymentRepo, penyewaRepo)

// Pass ke KamarService
kamarService := service.NewKamarService(kamarRepo, bookingRepo, paymentRepo, notificationService)
```

---

## Requirement Environment Variables

Tambahkan ke file `.env`:
```
# WhatsApp Configuration
FONNTE_TOKEN=your_fonnte_token_here
ADMIN_PHONE=6281234567890  # Nomor WhatsApp admin dengan kode negara
```

**Catatan:**
- `ADMIN_PHONE` format: gunakan kode negara (62 untuk Indonesia)
- Contoh: `6281234567890` (nomor 081234567890 dengan prefix 62)

---

## Database Query Generated

Query untuk mencari pending transfer payments by room ID:
```sql
SELECT * FROM pembayarans
INNER JOIN pemesanans ON pembayarans.pemesanan_id = pemesanans.id
WHERE pemesanans.kamar_id = ? 
  AND pembayarans.status_pembayaran = 'Pending'
  AND pembayarans.metode_pembayaran = 'transfer'
ORDER BY pembayarans.created_at DESC
```

---

## Error Handling

### Graceful Error Handling:
- Jika `ADMIN_PHONE` tidak dikonfigurasi: notifikasi tidak dikirim, kamar tetap terhapus
- Jika gagal mengirim WhatsApp: error di-log, kamar tetap terhapus, pemesanan tetap dibatalkan
- Jika gagal membatalkan pemesanan: error di-log, kamar tetap terhapus

### Non-Blocking Deletion:
- Penghapusan kamar tidak akan gagal meski notifikasi gagal dikirim
- Ini memastikan admin tetap bisa menghapus kamar, tetapi notifikasi yang penting tetap dikirim

---

## Testing Checklist

- [ ] Konfigurasi ADMIN_PHONE dan FONNTE_TOKEN di `.env`
- [ ] Buat pemesanan dengan metode transfer
- [ ] Upload bukti transfer (status pembayaran: Pending)
- [ ] Hapus kamar tersebut melalui admin panel
- [ ] Verifikasi:
  - [ ] Notifikasi WhatsApp diterima oleh admin
  - [ ] Pesan berisi nama, nomor HP, email penyewa
  - [ ] Pesan berisi jumlah pembayaran yang ditransfer
  - [ ] Pesan berisi link/path bukti transfer
  - [ ] Pemesanan berubah status menjadi "Cancelled"
  - [ ] Kamar berhasil dihapus

---

## Fitur Tambahan di Masa Depan

Bisa ditambahkan:
1. Email notifikasi ke penyewa tentang pembatalan pemesanan
2. Auto-refund process dengan tracking
3. History log untuk audit trail
4. Notifikasi SMS ke admin sebagai backup WhatsApp
5. Dashboard untuk admin melihat deleted rooms dengan pending payments

---

## Status Implementasi

✅ **SELESAI** - Semua fitur sudah diimplementasikan dan berhasil di-compile.

---

## File yang Diubah/Dibuat

### Dibuat:
- `internal/service/notification_service.go`

### Diubah:
- `internal/config/config.go`
- `internal/repository/payment_repository.go`
- `internal/service/kamar_service.go`
- `cmd/api/main.go`
- `.env.example`

---

## Catatan Penting

1. **Hanya untuk Transfer**: Fitur ini hanya berfungsi untuk pembayaran dengan metode transfer, bukan untuk pembayaran tunai/cash.
2. **Admin Notification**: Admin akan mendapat notifikasi WhatsApp dengan semua detail yang diperlukan untuk follow-up dengan penyewa.
3. **Booking Cancellation**: Pemesanan akan otomatis dibatalkan untuk mencegah konflik data.
4. **Non-Blocking**: Jika ada error saat mengirim notifikasi, kamar tetap bisa dihapus untuk mencegah blocking di admin panel.
