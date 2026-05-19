# API Specification - Report Admin Dashboard

## Overview
Dokumentasi lengkap untuk API endpoints yang digunakan oleh Report (Laporan) Admin Dashboard.

---

## 🎯 1. MAIN ENDPOINT - Get Dashboard Stats

### Endpoint
```
GET /api/dashboard/stats
```

### Authentication
- **Required**: Yes (JWT Token)
- **Role**: Admin only

### Query Parameters
```
(none - akan di-extend di masa depan)
```

### Response Format (200 OK)
```json
{
  "total_revenue": 16000000,
  "active_tenants": 5,
  "available_rooms": 1,
  "occupied_rooms": 1,
  "pending_payments": 6,
  "pending_revenue": 5000000,
  "rejected_payments": 2,
  "potential_revenue": 21000000,
  "monthly_trend": [
    {
      "month": "January",
      "revenue": 2000000
    },
    {
      "month": "February",
      "revenue": 2000000
    }
  ],
  "type_breakdown": [
    {
      "type": "Premium",
      "revenue": 8000000,
      "count": 2,
      "occupied": 1
    },
    {
      "type": "Standard",
      "revenue": 8000000,
      "count": 1,
      "occupied": 0
    }
  ],
  "demographics": [
    {
      "name": "18-25",
      "value": 5,
      "color": "#FF6B6B"
    },
    {
      "name": "26-35",
      "value": 3,
      "color": "#4ECDC4"
    }
  ],
  "recent_checkouts": [
    {
      "room_name": "101",
      "tenant_name": "Budi Santoso",
      "checkout_date": "2024-12-31T00:00:00Z",
      "reason": "Completed lease"
    }
  ]
}
```

### Field Explanations

| Field | Type | Description |
|-------|------|-------------|
| `total_revenue` | number | Total revenue dari pembayaran status "Confirmed" |
| `active_tenants` | number | Jumlah tenant dengan booking status "Aktif" |
| `available_rooms` | number | Jumlah kamar dengan status "Tersedia" |
| `occupied_rooms` | number | Jumlah kamar dengan tenant aktif |
| `pending_payments` | number | Jumlah pembayaran dengan status "Pending" |
| `pending_revenue` | number | Total jumlah dari pembayaran status "Pending" |
| `rejected_payments` | number | Jumlah pembayaran dengan status "Rejected" |
| `potential_revenue` | number | Total revenue (Confirmed + Pending) |
| `monthly_trend` | array | Monthly revenue breakdown (12 bulan terakhir) |
| `type_breakdown` | array | Revenue breakdown per room type |
| `demographics` | array | Age distribution tenant (calculated from tanggal_lahir) |
| `recent_checkouts` | array | Latest 5 checkouts |

### Error Responses

```json
// 401 Unauthorized
{
  "error": "missing or invalid token"
}

// 403 Forbidden (bukan admin)
{
  "error": "insufficient permissions"
}

// 500 Internal Server Error
{
  "error": "failed to fetch dashboard stats"
}
```

---

## 📊 2. GET ROOM OCCUPANCY DATA

### Endpoint
```
GET /api/dashboard/room-occupancy
```

### Response Format (200 OK)
```json
[
  {
    "room_id": 1,
    "nomor_kamar": "101",
    "tipe_kamar": "Premium",
    "status": "Aktif",
    "tenant_name": "Budi Santoso",
    "harga_per_bulan": 1500000,
    "is_paid_up_to_date": true,
    "checkout_date": "2024-12-31T00:00:00Z"
  },
  {
    "room_id": 2,
    "nomor_kamar": "102",
    "tipe_kamar": "Standard",
    "status": "Tersedia",
    "tenant_name": null,
    "harga_per_bulan": 1300000,
    "is_paid_up_to_date": null,
    "checkout_date": null
  }
]
```

---

## 👥 3. GET TENANT ROOMS DATA

### Endpoint
```
GET /api/dashboard/tenant-rooms
```

### Response Format (200 OK)
```json
[
  {
    "room_id": 1,
    "nomor_kamar": "101",
    "penyewa_id": 3,
    "tenant_name": "Budi Santoso",
    "nomor_hp": "081234567890",
    "email": "budi@email.com",
    "status": "Aktif",
    "check_in": "2024-01-15",
    "check_out": "2024-12-31",
    "durasi_sewa": 12
  }
]
```

---

## 💰 4. GET PAYMENTS BY ROOM

### Endpoint
```
GET /api/dashboard/payments/room/:id
```

### URL Parameters
- `id` (uint, required): Room ID

### Response Format (200 OK)
```json
{
  "room_id": 1,
  "nomor_kamar": "101",
  "tenant_name": "Budi Santoso",
  "penyewa_id": 3,
  "email": "budi@email.com",
  "nomor_hp": "081234567890",
  "check_in": "2024-01-15",
  "check_out": "2024-12-31",
  "durasi_sewa": 12,
  "payments": [
    {
      "id": 1,
      "jumlah_bayar": 2000000,
      "status_pembayaran": "Confirmed",
      "metode_pembayaran": "transfer",
      "tanggal_bayar": "2024-01-15",
      "payment_month": "January 2024",
      "bukti_transfer": "/uploads/proof_1.jpg"
    },
    {
      "id": 2,
      "jumlah_bayar": 2000000,
      "status_pembayaran": "Pending",
      "metode_pembayaran": "transfer",
      "tanggal_bayar": "2024-02-15",
      "payment_month": "February 2024",
      "bukti_transfer": "/uploads/proof_2.jpg"
    }
  ]
}
```

---

## 💰 5. GET PAYMENTS BY TENANT

### Endpoint
```
GET /api/dashboard/payments/tenant/:id
```

### URL Parameters
- `id` (uint, required): Penyewa (Tenant) ID

### Response Format (200 OK)
```json
{
  "nama_lengkap": "Budi Santoso",
  "email": "budi@email.com",
  "nomor_hp": "081234567890",
  "nik": "1234567890123456",
  "alamat_asal": "Jln. Merdeka No. 1",
  "jenis_kelamin": "M",
  "tanggal_lahir": "1995-05-10",
  "foto_profil": "/uploads/profile.jpg",
  "role": "tenant",
  "nomor_kamar": "101",
  "tipe_kamar": "Premium",
  "harga_per_bulan": 1500000,
  "check_in": "2024-01-15",
  "check_out": "2024-12-31",
  "durasi_sewa": 12,
  "payments": [
    {
      "id": 1,
      "jumlah_bayar": 2000000,
      "status_pembayaran": "Confirmed",
      "metode_pembayaran": "transfer",
      "tanggal_bayar": "2024-01-15",
      "payment_month": "January 2024",
      "bukti_transfer": "/uploads/proof_1.jpg"
    }
  ]
}
```

---

## 📋 6. GET ALL PAYMENTS

### Endpoint
```
GET /api/payments
```

### Query Parameters (Future Enhancement)
```
?status=Confirmed      # Filter by status
?metode=transfer       # Filter by payment method
?startDate=2024-01-01  # Filter by date range
&endDate=2024-12-31
?page=1&limit=100      # Pagination
```

### Response Format (200 OK)
```json
[
  {
    "id": 1,
    "pemesanan_id": 1,
    "jumlah_bayar": 2000000,
    "tanggal_bayar": "2024-01-15T10:30:00Z",
    "bukti_transfer": "/uploads/proof_1.jpg",
    "status_pembayaran": "Confirmed",
    "order_id": "ORDER-001",
    "metode_pembayaran": "transfer",
    "tipe_pembayaran": "full",
    "jumlah_dp": 0,
    "tanggal_jatuh_tempo": "2024-02-15T00:00:00Z",
    "idempotency_key": "key-001",
    "confirmed_at": "2024-01-15T10:35:00Z",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
]
```

---

## 🔧 FUTURE ENHANCEMENTS

### Date Range Filtering
```
GET /api/dashboard/stats?period=last_6_months
GET /api/dashboard/stats?startDate=2024-01-01&endDate=2024-06-30
GET /api/payments?dateRange=last_30_days

Supported periods:
- last_7_days
- last_30_days
- last_6_months
- last_12_months
- all_time
- custom (requires startDate & endDate)
```

### Pagination Support
```
GET /api/payments?page=1&limit=100
GET /api/dashboard/room-occupancy?page=1&limit=50

Response wrapper (future):
{
  "data": [...],
  "pagination": {
    "current_page": 1,
    "per_page": 100,
    "total": 500,
    "total_pages": 5
  }
}
```

### Field Selection (for bandwidth optimization)
```
GET /api/payments?fields=id,jumlah_bayar,status_pembayaran,tanggal_bayar
GET /api/dashboard/stats?exclude=demographics,recent_checkouts
```

---

## 🛠️ INTEGRATION NOTES

### Current Issues & Fixes

#### Issue: Total Revenue Label Stuck at "Rp 100.000"
**Cause**: Frontend parsing error atau response format mismatch

**Debug Checklist**:
1. Verify API returns valid JSON array/object
2. Check status value case: `Confirmed` vs `confirmed`
3. Verify `jumlah_bayar` is number not string
4. Check datetime format consistency

**Fix**:
```javascript
// Frontend parsing
const stats = response.data; // Should be object, not array
const confirmedPayments = stats.total_revenue; // Pre-calculated by backend
```

### Data Refresh Strategy

#### Recommended Caching:
- **Stat Cards**: Refresh every 30-60 seconds (not frequently changing)
- **Payments List**: Refresh every 60 seconds
- **Occupancy**: Refresh every 5 minutes
- **Demographics**: Cache until user manually refresh

#### Implementation:
```javascript
// React Hook Example
useEffect(() => {
  fetchStats();
  const interval = setInterval(fetchStats, 30000); // 30 seconds
  return () => clearInterval(interval);
}, []);
```

### Response Consistency Rules

1. **Status Values** (Case-Sensitive):
   - `Confirmed`, `Pending`, `Rejected`, `Cancelled`, `Failed`
   - NOT: `confirmed`, `pending`, etc.

2. **DateTime Format**:
   - Always ISO 8601: `2024-01-15T10:30:00Z`
   - NOT: `2024-01-15 10:30:00`

3. **Numeric Types**:
   - Currency as `number`, not `string`
   - Example: `1500000` NOT `"1500000"`

4. **Null Handling**:
   - Empty results: `[]` (not `null`)
   - Optional fields: `null` (not `undefined`)

---

## 📊 Data Calculation Examples

### Total Revenue Calculation
```
Total Revenue = SUM(jumlah_bayar) WHERE status_pembayaran = "Confirmed"
```

### Pending Revenue
```
Pending Revenue = SUM(jumlah_bayar) WHERE status_pembayaran = "Pending" AND metode_pembayaran = "transfer"
```

### Average Room Price
```
Average Price = SUM(harga_per_bulan) / COUNT(kamar) WHERE deleted_at IS NULL
```

### Occupancy Rate
```
Occupancy Rate = (occupied_rooms / total_rooms) * 100
```

---

## 🔐 Security Notes

- All endpoints require admin authentication
- Do not expose `hidden_reason`, `bank_account`, or sensitive data
- Rate limit: 100 requests per minute per user
- Sensitive data in responses should be hashed or masked

---

## 📈 Performance Considerations

### Typical Response Sizes
- `/api/dashboard/stats` → 5-15 KB
- `/api/payments` (full) → 50-200 KB
- `/api/dashboard/room-occupancy` → 2-5 KB

### Optimization Tips
1. Use pagination for payment lists
2. Cache stats response client-side
3. Consider field filtering for large responses
4. Lazy load demographics data

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2024-12-06 | Initial specification |
| (Future) | TBD | Date range filtering |
| (Future) | TBD | Pagination support |
| (Future) | TBD | Real-time updates via WebSocket |
