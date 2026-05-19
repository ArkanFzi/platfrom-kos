# Frontend Integration Guide - Report Admin Dashboard

## Overview
Panduan lengkap untuk mengintegrasikan Report Admin Dashboard dengan API backend.

---

## 🚀 QUICK START

### 1. Main Stats Hook
```typescript
// hooks/useReportStats.ts
import { useEffect, useState } from 'react';
import axios from 'axios';

interface DashboardStats {
  total_revenue: number;
  active_tenants: number;
  available_rooms: number;
  occupied_rooms: number;
  pending_payments: number;
  pending_revenue: number;
  rejected_payments: number;
  potential_revenue: number;
  monthly_trend: MonthlyData[];
  type_breakdown: TypeRevenue[];
  demographics: Demographic[];
  recent_checkouts: RecentCheckout[];
}

export const useReportStats = () => {
  const [stats, setStats] = useState<DashboardStats | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchStats = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.get('/api/dashboard/stats');
      // Response is already an object, not an array!
      setStats(response.data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchStats();
    // Auto-refresh every 30 seconds
    const interval = setInterval(fetchStats, 30000);
    return () => clearInterval(interval);
  }, []);

  return { stats, loading, error, refetch: fetchStats };
};
```

### 2. Stat Cards Component
```typescript
// components/StatCard.tsx
import React from 'react';

interface StatCardProps {
  label: string;
  value: number;
  format: 'currency' | 'number' | 'percentage';
  trend?: number;
  icon?: React.ReactNode;
}

export const StatCard: React.FC<StatCardProps> = ({
  label,
  value,
  format,
  trend,
  icon,
}) => {
  const formatValue = () => {
    switch (format) {
      case 'currency':
        return `Rp ${new Intl.NumberFormat('id-ID').format(Math.round(value))}`;
      case 'percentage':
        return `${value.toFixed(1)}%`;
      default:
        return value.toString();
    }
  };

  return (
    <div className="stat-card">
      <div className="stat-label">{label}</div>
      <div className="stat-value">{formatValue()}</div>
      {trend !== undefined && (
        <div className={`stat-trend ${trend >= 0 ? 'positive' : 'negative'}`}>
          {trend >= 0 ? '↑' : '↓'} {Math.abs(trend)}%
        </div>
      )}
    </div>
  );
};
```

### 3. Report Dashboard Page
```typescript
// pages/admin/report.tsx
import React, { useMemo } from 'react';
import { useReportStats } from '@/hooks/useReportStats';
import { StatCard } from '@/components/StatCard';
import { RevenueChart } from '@/components/RevenueChart';
import { OccupancyChart } from '@/components/OccupancyChart';
import { DemographicsChart } from '@/components/DemographicsChart';

export default function ReportPage() {
  const { stats, loading, error } = useReportStats();

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;
  if (!stats) return <div>No data</div>;

  // Calculate occupancy percentage
  const occupancyRate = useMemo(() => {
    const total = stats.occupied_rooms + stats.available_rooms;
    if (total === 0) return 0;
    return (stats.occupied_rooms / total) * 100;
  }, [stats]);

  return (
    <div className="report-dashboard">
      {/* Stat Cards */}
      <div className="stat-cards-grid">
        <StatCard
          label="Total Revenue"
          value={stats.total_revenue}
          format="currency"
        />
        <StatCard
          label="Pending Revenue"
          value={stats.pending_revenue}
          format="currency"
        />
        <StatCard
          label="Average Rate"
          value={stats.type_breakdown.reduce((sum, t) => sum + t.revenue, 0) / 
                  stats.type_breakdown.length}
          format="currency"
        />
        <StatCard
          label="Occupancy Rate"
          value={occupancyRate}
          format="percentage"
        />
      </div>

      {/* Charts */}
      <div className="charts-grid">
        <RevenueChart data={stats.type_breakdown} />
        <OccupancyChart
          occupied={stats.occupied_rooms}
          available={stats.available_rooms}
        />
        <DemographicsChart data={stats.demographics} />
      </div>

      {/* Recent Checkouts */}
      <div className="recent-checkouts">
        <h3>Recent Checkouts</h3>
        <ul>
          {stats.recent_checkouts.map((checkout) => (
            <li key={`${checkout.room_name}-${checkout.checkout_date}`}>
              <strong>{checkout.room_name}</strong> - {checkout.tenant_name}
              <small>{new Date(checkout.checkout_date).toLocaleDateString()}</small>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
```

---

## 🐛 DEBUGGING TOTAL REVENUE ISSUE

### Issue: Total Revenue Stuck at "Rp 100.000"

#### Root Cause Analysis

**Scenario 1: Frontend Parsing Error**
```typescript
// ❌ WRONG - Treating array as single value
const totalRevenue = response.data[0]; // This gets first payment item, not stats!

// ✅ CORRECT - Stats is object, not array
const totalRevenue = response.data.total_revenue;
```

**Scenario 2: Case Sensitivity in Status Filter**
```typescript
// ❌ WRONG - Status is "Confirmed" not "confirmed"
const confirmedPayments = payments.filter(p => p.status_pembayaran === 'confirmed');

// ✅ CORRECT
const confirmedPayments = payments.filter(p => p.status_pembayaran === 'Confirmed');
```

**Scenario 3: Type Coercion Issue**
```typescript
// ❌ WRONG - jumlah_bayar comes as string
const total = payments.reduce((sum, p) => sum + p.jumlah_bayar, 0);
// Result: "000" + 1000000 = "0001000000" (string concatenation!)

// ✅ CORRECT - Ensure numeric type
const total = payments.reduce(
  (sum, p) => sum + parseFloat(p.jumlah_bayar),
  0
);
```

**Scenario 4: Backend Response Changed**
```typescript
// OLD Response (assuming payments array)
[
  { id: 1, jumlah_bayar: 2000000, status_pembayaran: "Confirmed" }
]

// NEW Response (stats object - what backend returns now)
{
  total_revenue: 16000000,
  pending_revenue: 5000000,
  ...
}
```

### Debug Steps

1. **Check API Response in Browser DevTools**
   ```javascript
   // Open Console, run:
   fetch('/api/dashboard/stats')
     .then(r => r.json())
     .then(data => console.log(JSON.stringify(data, null, 2)));
   ```

2. **Verify Data Type**
   ```javascript
   console.log(typeof stats.total_revenue); // Should be 'number', not 'string'
   console.log(Array.isArray(stats)); // Should be false
   ```

3. **Check Status Values**
   ```javascript
   console.log(stats); // Look for exact values
   // Should see: "Confirmed", "Pending", "Rejected"
   // NOT: "confirmed", "pending", "rejected"
   ```

4. **Validate DateTime Format**
   ```javascript
   const date = new Date(stats.monthly_trend[0].date);
   console.log(date.toISOString()); // Should parse correctly
   ```

### Quick Fix Checklist

- [ ] Verify API endpoint is `/api/dashboard/stats` (not `/api/payments`)
- [ ] Check response is object `{}` not array `[]`
- [ ] Ensure `total_revenue` field exists in response
- [ ] Verify value is `number` type
- [ ] Test with fresh API call (clear cache)
- [ ] Check browser console for parsing errors
- [ ] Validate JWT token is still valid

---

## 📊 CHART INTEGRATION

### Revenue by Room Type Chart
```typescript
// components/RevenueChart.tsx
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Legend } from 'recharts';

interface ChartProps {
  data: TypeRevenue[];
}

export const RevenueChart: React.FC<ChartProps> = ({ data }) => {
  const chartData = data.map(item => ({
    name: item.type,
    revenue: item.revenue / 1000000, // Convert to millions for readability
    occupied: item.occupied,
  }));

  return (
    <BarChart width={600} height={300} data={chartData}>
      <CartesianGrid strokeDasharray="3 3" />
      <XAxis dataKey="name" />
      <YAxis label={{ value: 'Revenue (Million Rp)', angle: -90 }} />
      <Legend />
      <Bar dataKey="revenue" fill="#82ca9d" name="Revenue (M Rp)" />
      <Bar dataKey="occupied" fill="#ffc658" name="Occupied Rooms" />
    </BarChart>
  );
};
```

### Occupancy Donut Chart
```typescript
// components/OccupancyChart.tsx
import { PieChart, Pie, Cell, Legend } from 'recharts';

interface OccupancyChartProps {
  occupied: number;
  available: number;
}

export const OccupancyChart: React.FC<OccupancyChartProps> = ({
  occupied,
  available,
}) => {
  const data = [
    { name: 'Occupied', value: occupied, fill: '#FF6B6B' },
    { name: 'Available', value: available, fill: '#4ECDC4' },
  ];

  return (
    <PieChart width={400} height={300}>
      <Pie
        data={data}
        cx={200}
        cy={150}
        innerRadius={60}
        outerRadius={100}
        paddingAngle={5}
        dataKey="value"
      >
        {data.map((entry, index) => (
          <Cell key={`cell-${index}`} fill={entry.fill} />
        ))}
      </Pie>
      <Legend />
    </PieChart>
  );
};
```

### Demographics Age Distribution Chart
```typescript
// components/DemographicsChart.tsx
import { BarChart, Bar, XAxis, YAxis, CartesianGrid } from 'recharts';

interface DemographicsChartProps {
  data: Demographic[];
}

export const DemographicsChart: React.FC<DemographicsChartProps> = ({
  data,
}) => {
  return (
    <BarChart width={600} height={300} data={data}>
      <CartesianGrid strokeDasharray="3 3" />
      <XAxis dataKey="name" label={{ value: 'Age Group', position: 'insideBottom' }} />
      <YAxis label={{ value: 'Tenant Count', angle: -90 }} />
      <Bar dataKey="value" fill="#8884d8" />
    </BarChart>
  );
};
```

---

## 🔄 DATE RANGE FILTERING (Future Feature)

### Implementation
```typescript
// hooks/useFilteredStats.ts
interface FilterOptions {
  period?: 'last_7_days' | 'last_30_days' | 'last_6_months' | 'last_12_months' | 'all_time';
  startDate?: string;
  endDate?: string;
}

export const useFilteredStats = (filters: FilterOptions) => {
  const [stats, setStats] = useState<DashboardStats | null>(null);

  useEffect(() => {
    const params = new URLSearchParams();
    if (filters.period) params.append('period', filters.period);
    if (filters.startDate) params.append('startDate', filters.startDate);
    if (filters.endDate) params.append('endDate', filters.endDate);

    axios.get(`/api/dashboard/stats?${params}`).then(res => {
      setStats(res.data);
    });
  }, [filters]);

  return stats;
};
```

### UI Component
```typescript
<div className="filter-controls">
  <button onClick={() => setPeriod('last_7_days')}>Last 7 Days</button>
  <button onClick={() => setPeriod('last_30_days')}>Last 30 Days</button>
  <button onClick={() => setPeriod('last_6_months')}>Last 6 Months</button>
  <button onClick={() => setPeriod('last_12_months')}>Last Year</button>
  <button onClick={() => setPeriod('all_time')}>All Time</button>
</div>
```

---

## 💾 CACHING STRATEGY

### Local Storage + Expiry
```typescript
const CACHE_DURATION = 30 * 60 * 1000; // 30 minutes

const getCachedStats = () => {
  const cached = localStorage.getItem('dashboard_stats');
  const timestamp = localStorage.getItem('dashboard_stats_time');

  if (!cached || !timestamp) return null;

  const age = Date.now() - parseInt(timestamp);
  if (age > CACHE_DURATION) {
    localStorage.removeItem('dashboard_stats');
    localStorage.removeItem('dashboard_stats_time');
    return null;
  }

  return JSON.parse(cached);
};

const setCachedStats = (stats: DashboardStats) => {
  localStorage.setItem('dashboard_stats', JSON.stringify(stats));
  localStorage.setItem('dashboard_stats_time', Date.now().toString());
};
```

---

## ⚠️ ERROR HANDLING

```typescript
const fetchStats = async () => {
  try {
    const response = await axios.get('/api/dashboard/stats');
    
    // Validate response structure
    if (!response.data || typeof response.data !== 'object') {
      throw new Error('Invalid response format');
    }
    
    // Validate required fields
    const requiredFields = [
      'total_revenue',
      'active_tenants',
      'available_rooms',
    ];
    
    for (const field of requiredFields) {
      if (!(field in response.data)) {
        throw new Error(`Missing required field: ${field}`);
      }
    }
    
    setStats(response.data);
  } catch (error) {
    if (error.response?.status === 401) {
      // Redirect to login
      window.location.href = '/login';
    } else if (error.response?.status === 403) {
      setError('Admin access required');
    } else {
      setError(error.message);
    }
  }
};
```

---

## 📱 RESPONSIVE DESIGN

```css
@media (max-width: 768px) {
  .stat-cards-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .charts-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 480px) {
  .stat-cards-grid {
    grid-template-columns: 1fr;
  }
}
```

---

## 🚨 CURRENT FIX FOR TOTAL REVENUE ISSUE

Replace your current Total Revenue fetching code with:

```typescript
// BEFORE (Wrong)
const [totalRevenue, setTotalRevenue] = useState('Rp 100.000');
// ... somehow parsing only 1 payment

// AFTER (Correct)
const { stats } = useReportStats();
const totalRevenue = stats?.total_revenue || 0;

return (
  <StatCard
    label="Total Revenue"
    value={totalRevenue}
    format="currency"
  />
);
```

This uses the backend's pre-calculated `total_revenue` field instead of trying to parse and sum payments on the frontend.
