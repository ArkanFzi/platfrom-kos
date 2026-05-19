# FAQ - Report Admin API Integration

Jawaban detail untuk semua pertanyaan tentang API Report Admin Dashboard.

---

## STAT CARDS

### Q1: Endpoint mana untuk Total Revenue?
**A:** 
```
GET /api/dashboard/stats
```
Field: `total_revenue`

Response SUDAH pre-calculated oleh backend. Jangan parsing payments manual, ambil `total_revenue` langsung.

---

### Q2: Apakah Total Revenue bisa di-filter by date range?
**A:** 
**Currently:** ❌ Tidak support (akan di-add di v2)

**Future:** ✅ akan support
```
GET /api/dashboard/stats?period=last_6_months
GET /api/dashboard/stats?startDate=2024-01-01&endDate=2024-06-30
```

---

### Q3: Apakah Total Revenue hanya menghitung status "Confirmed"?
**A:** 
**Benar** ✅

```
total_revenue = SUM(jumlah_bayar) WHERE status_pembayaran = "Confirmed"
```

Untuk pending revenue, gunakan field `pending_revenue` (sudah terpisah).

---

### Q4: Format response apa? Direct calculation atau array payments?
**A:** 
**Direct calculation** ✅ (option 1 yang Anda minta)

Response:
```json
{
  "total_revenue": 16000000,
  "pending_revenue": 5000000,
  ...
}
```

BUKAN array payments. Backend sudah sum-kan untuk Anda.

---

### Q5: Ada max data limit atau pagination?
**A:** 
- Current: ❌ Tidak ada
- Future: ✅ Akan di-add
  ```
  GET /api/dashboard/stats?page=1&limit=100
  ```

---

## PENDING REVENUE

### Q6: Apakah Pending Revenue menggunakan endpoint sama?
**A:** 
**Ya** ✅

Endpoint yang sama: `/api/dashboard/stats`

Field: `pending_revenue`

---

### Q7: Apakah Pending Revenue hanya status "Pending" atau ada filter lain?
**A:** 
```
pending_revenue = SUM(jumlah_bayar) WHERE 
  status_pembayaran = "Pending" 
  AND metode_pembayaran = "transfer"
```

Hanya yang metode transfer (tidak cash/tunai).

---

### Q8: Apakah pending revenue dari JumlahBayar atau JumlahDP?
**A:** 
**JumlahBayar** ✅

Field: `jumlah_bayar` (bukan `jumlah_dp`)

---

## AVERAGE RATE

### Q9: Endpoint untuk Average Room Price?
**A:** 
Included dalam:
```
GET /api/dashboard/stats
```

Field: `type_breakdown` (array dengan revenue per tipe kamar)

Contoh calculation:
```
Average = SUM(type.revenue) / COUNT(type)
```

---

### Q10: Average dari semua kamar atau hanya kamar terisi?
**A:** 
**Dari semua kamar** (available + occupied)

```
Average Price = SUM(harga_per_bulan) / COUNT(*) 
WHERE deleted_at IS NULL
```

---

### Q11: Format response: Direct calculation atau array rooms?
**A:** 
**Array** ✅ (grouped by tipe kamar)

```json
{
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
  ]
}
```

Frontend bisa calculate average dari array ini.

---

## OCCUPANCY RATE

### Q12: Endpoint untuk Occupancy Rate?
**A:** 
```
GET /api/dashboard/stats
```

Fields: 
- `occupied_rooms` 
- `available_rooms`

Calculation: `(occupied / total) * 100`

---

### Q13: Status kamar apa aja yang possible?
**A:** 
```
"Tersedia"  - Empty/Available
"Penuh"     - Occupied
"Maintenance" - Under maintenance (jarang digunakan)
```

Backend hitung occupied dari booking status, bukan room status langsung.

---

### Q14: Format response untuk occupancy?
**A:** 
```json
{
  "occupied_rooms": 1,
  "available_rooms": 1,
  "type_breakdown": [
    {
      "type": "Premium",
      "count": 2,
      "occupied": 1
    }
  ]
}
```

Frontend calculate percentage: `(1 / 2) * 100 = 50%`

---

## GRAFIK - REVENUE BY ROOM TYPE

### Q15: Endpoint untuk Revenue Breakdown per Room Type?
**A:** 
Included dalam `/api/dashboard/stats`:
```json
"type_breakdown": [
  {
    "type": "Premium",
    "revenue": 8000000,
    "count": 2,
    "occupied": 1
  }
]
```

---

### Q16: Data sudah aggregated atau array perlu di-group?
**A:** 
**Sudah aggregated** ✅ oleh backend

Array already grouped by room type dan pre-summed.

---

### Q17: Period filter supported?
**A:** 
- Current: ❌ No
- Future: ✅ Will support
  ```
  GET /api/dashboard/stats?period=last_6_months
  ```

---

## GRAFIK - TENANT DEMOGRAPHICS

### Q18: Endpoint untuk Tenant Demographics?
**A:** 
```
GET /api/dashboard/stats
```

Field: `demographics` (array)

```json
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
]
```

---

### Q19: Ada field "age" atau "birthDate" di tenant model?
**A:** 
**birthDate** ✅

Field: `tanggal_lahir` (format: `YYYY-MM-DD`)

Backend hitung age dari `tanggal_lahir`.

---

### Q20: Bagaimana age grouping?
**A:** 
```
18-25
26-35
36-45
46-55
56+
```

Backend sudah group-kan. Frontend tinggal display.

---

## DATE RANGE FILTERING

### Q21: Fetch data baru atau filter client-side?
**A:** 
**Fetch data baru** ✅ (recommended)

API akan di-extend untuk support:
```
GET /api/dashboard/stats?period=last_30_days
GET /api/dashboard/stats?startDate=...&endDate=...
```

---

### Q22: Query parameter format apa?
**A:** 
Recommended format:
```
GET /api/dashboard/stats?period=last_6_months

Supported periods:
- last_7_days
- last_30_days
- last_6_months
- last_12_months
- all_time
- custom (+ startDate & endDate)
```

---

## DATA CONSISTENCY & CACHING

### Q23: Berapa lama data seharusnya di-cache?
**A:** 
Recommendations:
- **Stat Cards**: 30-60 seconds (jarang berubah)
- **Revenue Charts**: 60 seconds
- **Occupancy**: 5 minutes
- **Demographics**: 30 minutes

---

### Q24: Ada webhook atau real-time update?
**A:** 
- Current: ❌ No
- Future: ✅ WebSocket support planned

---

### Q25: Data berubah setiap berapa detik/menit?
**A:** 
Typical frequency:
- **Revenue**: Change saat ada payment confirmation (unpredictable)
- **Occupancy**: Change saat booking/checkout (unpredictable)
- **Demographics**: Change saat tenant register (unpredictable)

Safe cache 15 detik: ✅ OK

---

### Q26: Safe untuk cache dan update setiap 15 detik?
**A:** 
**Ya** ✅ Recommended

15 detik adalah sweet spot antara:
- Fresh data (terlihat real-time untuk user)
- Not hammering backend
- Battery-friendly untuk mobile

---

## ERROR HANDLING & PAGINATION

### Q27: Fallback data jika API error?
**A:** 
Implement graceful fallback:
```typescript
try {
  const stats = await fetchStats();
  setStats(stats);
} catch (error) {
  // Use cached stats if available
  const cached = getCachedStats();
  if (cached) {
    setStats(cached);
    showWarning('Showing cached data from 5 min ago');
  } else {
    setError('Failed to load report');
  }
}
```

---

### Q28: Ada pagination untuk large dataset?
**A:** 
- Current: ❌ No (tapi data terbatas, tidak besar)
- Future: ✅ Support pagination
  ```
  GET /api/payments?page=1&limit=100
  ```

---

## CURRENT ISSUE - DEBUG

### Q29: Bagaimana format JSON response dari /api/payments?
**A:** 
Array of payments (berbeda dari /api/dashboard/stats):
```json
[
  {
    "id": 1,
    "pemesanan_id": 1,
    "jumlah_bayar": 2000000,
    "status_pembayaran": "Confirmed",
    "metode_pembayaran": "transfer",
    "tanggal_bayar": "2024-01-15T10:30:00Z",
    ...
  },
  ...
]
```

**Catatan:** `/api/dashboard/stats` return OBJECT, bukan array!

---

### Q30: Apakah struktur response sesuai model?
**A:** 
**Ya** ✅ Sesuai model Pembayaran:
- ✅ `id`
- ✅ `pemesanan_id`
- ✅ `jumlah_bayar` (number)
- ✅ `status_pembayaran` (string: "Confirmed", "Pending", "Rejected")
- ✅ `metode_pembayaran`
- ✅ `tanggal_bayar` (ISO 8601 datetime)
- ✅ `bukti_transfer` (URL string)

---

### Q31: Ada case sensitivity issue?
**A:** 
**Ya** ⚠️ **Case-sensitive!**

Status values:
```
"Confirmed"  ✅ (capital C)
"Pending"    ✅ (capital P)
"Rejected"   ✅ (capital R)
"Cancelled"  ✅ (capital C)

NOT: "confirmed", "pending", "rejected", "cancelled"
```

**Fix:**
```typescript
// ❌ Wrong
payments.filter(p => p.status_pembayaran === 'confirmed')

// ✅ Correct
payments.filter(p => p.status_pembayaran === 'Confirmed')
```

---

## BANDWIDTH & PERFORMANCE

### Q32: Typical response size?
**A:** 
- `/api/dashboard/stats` → 5-15 KB
- `/api/payments` (full) → 50-200 KB (depending on record count)
- `/api/dashboard/room-occupancy` → 2-5 KB

---

### Q33: Exclude fields untuk reduce payload?
**A:** 
Future feature (planned v2):
```
GET /api/payments?fields=id,jumlah_bayar,status_pembayaran
GET /api/dashboard/stats?exclude=demographics,recent_checkouts
```

Currently, tidak support query-based field filtering.

---

## RECOMMENDATIONS

### For Optimal Performance:

1. **Cache aggressively**
   - Stat cards: 30s
   - Payments: 60s
   - Use localStorage

2. **Use the right endpoint**
   - Dashboard stats: `/api/dashboard/stats`
   - Detailed payments: `/api/payments`
   - By room: `/api/dashboard/payments/room/:id`
   - By tenant: `/api/dashboard/payments/tenant/:id`

3. **Handle errors gracefully**
   - Fallback to cached data
   - Show user-friendly messages
   - Log errors for debugging

4. **Validate data structure**
   - Check response type (object vs array)
   - Verify required fields exist
   - Handle null/undefined values

5. **DateTime handling**
   - Always use ISO 8601 format
   - Parse with `new Date(dateString)`
   - Format for display: `date.toLocaleDateString('id-ID')`

---

## TEMPLATE UNTUK BACKEND TEAM

```markdown
Halo Backend Team,

Ini adalah dokumentasi API Report Admin yang kami compile dari analisis code.
Beberapa clarification & confirmation needed:

1. ✅ `/api/dashboard/stats` return object dengan pre-calculated fields → CORRECT?
2. ✅ Status values case-sensitive ("Confirmed", bukan "confirmed") → CORRECT?
3. ✅ `total_revenue` only includes "Confirmed" payments → CORRECT?
4. ❓ Future: Support date range filtering (period param)?
5. ❓ Future: Support pagination untuk payments list?
6. ❓ Future: Support field filtering untuk reduce payload?

Current response structure terlihat lengkap dan well-designed. 
Frontend team bisa mulai integrate dengan confidence menggunakan docs ini.

Thanks!
```

---

**Dokumen dibuat:** 2024-12-06  
**Version:** 1.0  
**Last Updated:** 2024-12-06
