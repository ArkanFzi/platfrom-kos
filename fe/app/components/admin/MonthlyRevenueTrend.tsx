import React, { useEffect, useState } from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import axios from "axios";

interface RevenueData {
  month: string;
  revenue: number;
}

export default function MonthlyRevenueTrend() {
  const [data, setData] = useState<RevenueData[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    async function fetchRevenue() {
      setLoading(true);
      setError("");
      try {
        // Ganti URL berikut dengan endpoint backend yang sesuai
        const res = await axios.get("/api/revenue/monthly?months=6");
        setData(res.data);
      } catch (err: any) {
        setError("Gagal memuat data revenue");
      } finally {
        setLoading(false);
      }
    }
    fetchRevenue();
  }, []);

  if (loading) return <div>Loading revenue trend...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div style={{ width: "100%", height: 320 }}>
      <h3 className="text-lg font-bold mb-2">
        Monthly Revenue Trend (Last 6 Months)
      </h3>
      <ResponsiveContainer width="100%" height={280}>
        <LineChart
          data={data}
          margin={{ top: 20, right: 30, left: 0, bottom: 0 }}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="month" />
          <YAxis tickFormatter={(v) => `Rp${v.toLocaleString()}`} />
          <Tooltip formatter={(v) => `Rp${Number(v).toLocaleString()}`} />
          <Line
            type="monotone"
            dataKey="revenue"
            stroke="#8884d8"
            strokeWidth={3}
            dot={{ r: 5 }}
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}
