import React, { useState, useMemo } from 'react';
import { motion } from 'framer-motion';
import { ChevronLeft, ChevronRight, Home, DollarSign, Clock, TrendingUp } from 'lucide-react';
import { Card } from '@/app/components/ui/card';
import { Button } from '@/app/components/ui/button';

interface BookingStats {
  totalBookings: number;
  activeBookings: number;
  completedBookings: number;
  totalSpent: number;
  nextPaymentDate: string;
  upcomingPayments: number;
}

interface SmartCalendarProps {
  stats?: BookingStats;
  showStats?: boolean;
}

export function SmartCalendar({ stats, showStats = false }: SmartCalendarProps) {
  const [currentDate, setCurrentDate] = useState(new Date(2026, 1, 4)); // Feb 4, 2026

  const monthYear = currentDate.toLocaleDateString('id-ID', { month: 'long', year: 'numeric' });
  const daysInMonth = new Date(currentDate.getFullYear(), currentDate.getMonth() + 1, 0).getDate();
  const firstDayOfMonth = new Date(currentDate.getFullYear(), currentDate.getMonth(), 1).getDay();

  const days = Array.from({ length: daysInMonth }, (_, i) => i + 1);
  const emptyDays = Array.from({ length: firstDayOfMonth }, (_, i) => i);

  // Mock booking events
  const bookingEvents = useMemo(() => {
    return [
      { date: 10, status: 'active', label: 'Move-in: Apt 301' },
      { date: 15, status: 'payment', label: 'Payment Due' },
      { date: 20, status: 'reminder', label: 'Next Payment' },
      { date: 28, status: 'active', label: 'Move-in: Studio B' },
    ];
  }, []);

  const getEventForDay = (day: number) => bookingEvents.find(e => e.date === day);

  const prevMonth = () => {
    setCurrentDate(new Date(currentDate.getFullYear(), currentDate.getMonth() - 1));
  };

  const nextMonth = () => {
    setCurrentDate(new Date(currentDate.getFullYear(), currentDate.getMonth() + 1));
  };

  const eventColors = {
    active: 'bg-emerald-100 dark:bg-emerald-900/30 border-emerald-300 dark:border-emerald-800',
    payment: 'bg-red-100 dark:bg-red-900/30 border-red-300 dark:border-red-800',
    reminder: 'bg-orange-100 dark:bg-orange-900/30 border-orange-300 dark:border-orange-800',
  };

  const eventTextColors = {
    active: 'text-emerald-700 dark:text-emerald-300',
    payment: 'text-red-700 dark:text-red-300',
    reminder: 'text-orange-700 dark:text-orange-300',
  };

  return (
    <div className="w-full space-y-6">
      {/* Stats Cards */}
      {showStats && stats && (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          {/* Active Bookings */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.1 }}
          >
            <Card className="p-6 bg-gradient-to-br from-emerald-50 to-emerald-100 dark:from-emerald-950/50 dark:to-emerald-900/50 border-emerald-200 dark:border-emerald-800">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm text-emerald-700 dark:text-emerald-300 font-medium">Active Bookings</p>
                  <p className="text-3xl font-bold text-emerald-900 dark:text-emerald-100 mt-2">{stats.activeBookings}</p>
                </div>
                <div className="p-3 bg-emerald-200/50 dark:bg-emerald-800/50 rounded-lg">
                  <Home className="w-6 h-6 text-emerald-700 dark:text-emerald-300" />
                </div>
              </div>
            </Card>
          </motion.div>

          {/* Total Spent */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.2 }}
          >
            <Card className="p-6 bg-gradient-to-br from-blue-50 to-blue-100 dark:from-blue-950/50 dark:to-blue-900/50 border-blue-200 dark:border-blue-800">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm text-blue-700 dark:text-blue-300 font-medium">Total Spent</p>
                  <p className="text-3xl font-bold text-blue-900 dark:text-blue-100 mt-2">${stats.totalSpent.toLocaleString()}</p>
                </div>
                <div className="p-3 bg-blue-200/50 dark:bg-blue-800/50 rounded-lg">
                  <DollarSign className="w-6 h-6 text-blue-700 dark:text-blue-300" />
                </div>
              </div>
            </Card>
          </motion.div>

          {/* Next Payment */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.3 }}
          >
            <Card className="p-6 bg-gradient-to-br from-orange-50 to-orange-100 dark:from-orange-950/50 dark:to-orange-900/50 border-orange-200 dark:border-orange-800">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm text-orange-700 dark:text-orange-300 font-medium">Next Payment</p>
                  <p className="text-lg font-bold text-orange-900 dark:text-orange-100 mt-2">{stats.nextPaymentDate}</p>
                </div>
                <div className="p-3 bg-orange-200/50 dark:bg-orange-800/50 rounded-lg">
                  <Clock className="w-6 h-6 text-orange-700 dark:text-orange-300" />
                </div>
              </div>
            </Card>
          </motion.div>

          {/* Completed Bookings */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.4 }}
          >
            <Card className="p-6 bg-gradient-to-br from-purple-50 to-purple-100 dark:from-purple-950/50 dark:to-purple-900/50 border-purple-200 dark:border-purple-800">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm text-purple-700 dark:text-purple-300 font-medium">Completed</p>
                  <p className="text-3xl font-bold text-purple-900 dark:text-purple-100 mt-2">{stats.completedBookings}</p>
                </div>
                <div className="p-3 bg-purple-200/50 dark:bg-purple-800/50 rounded-lg">
                  <TrendingUp className="w-6 h-6 text-purple-700 dark:text-purple-300" />
                </div>
              </div>
            </Card>
          </motion.div>
        </div>
      )}

      {/* Calendar */}
      <Card className="p-6 bg-white dark:bg-slate-900 shadow-lg">
        {/* Header */}
        <div className="flex items-center justify-between mb-8">
          <h2 className="text-2xl font-bold text-slate-900 dark:text-white">{monthYear}</h2>
          <div className="flex gap-2">
            <Button
              variant="outline"
              size="sm"
              onClick={prevMonth}
              className="border-slate-200 dark:border-slate-700 hover:bg-slate-100 dark:hover:bg-slate-800"
            >
              <ChevronLeft className="w-4 h-4" />
            </Button>
            <Button
              variant="outline"
              size="sm"
              onClick={nextMonth}
              className="border-slate-200 dark:border-slate-700 hover:bg-slate-100 dark:hover:bg-slate-800"
            >
              <ChevronRight className="w-4 h-4" />
            </Button>
          </div>
        </div>

        {/* Weekday headers */}
        <div className="grid grid-cols-7 gap-2 mb-4">
          {['Min', 'Sen', 'Sel', 'Rab', 'Kam', 'Jum', 'Sab'].map((day) => (
            <div key={day} className="text-center font-semibold text-slate-600 dark:text-slate-400 text-sm py-2">
              {day}
            </div>
          ))}
        </div>

        {/* Calendar Grid */}
        <div className="grid grid-cols-7 gap-2">
          {/* Empty cells */}
          {emptyDays.map((i) => (
            <div key={`empty-${i}`} className="aspect-square" />
          ))}

          {/* Days */}
          {days.map((day) => {
            const event = getEventForDay(day);
            const isToday = day === 4 && currentDate.getMonth() === 1;

            return (
              <motion.div
                key={day}
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
                className={`aspect-square p-2 rounded-lg border-2 cursor-pointer transition-all ${
                  event
                    ? `${eventColors[event.status as keyof typeof eventColors]} border-current`
                    : isToday
                    ? 'bg-slate-100 dark:bg-slate-800 border-slate-300 dark:border-slate-600'
                    : 'bg-slate-50 dark:bg-slate-800/50 border-slate-200 dark:border-slate-700 hover:border-slate-400 dark:hover:border-slate-500'
                }`}
              >
                <p className={`text-sm font-semibold ${
                  event
                    ? eventTextColors[event.status as keyof typeof eventTextColors]
                    : isToday
                    ? 'text-slate-900 dark:text-white'
                    : 'text-slate-600 dark:text-slate-400'
                }`}>
                  {day}
                </p>
                {event && (
                  <p className={`text-xs mt-1 truncate font-medium ${
                    eventTextColors[event.status as keyof typeof eventTextColors]
                  }`}>
                    {event.label}
                  </p>
                )}
              </motion.div>
            );
          })}
        </div>
      </Card>

      {/* Legend */}
      <Card className="p-4 bg-slate-50 dark:bg-slate-800/50">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div className="flex items-center gap-3">
            <div className="w-3 h-3 rounded-full bg-emerald-500" />
            <span className="text-sm text-slate-700 dark:text-slate-300">Booking Aktif</span>
          </div>
          <div className="flex items-center gap-3">
            <div className="w-3 h-3 rounded-full bg-red-500" />
            <span className="text-sm text-slate-700 dark:text-slate-300">Pembayaran Jatuh Tempo</span>
          </div>
          <div className="flex items-center gap-3">
            <div className="w-3 h-3 rounded-full bg-orange-500" />
            <span className="text-sm text-slate-700 dark:text-slate-300">Pengingat Pembayaran</span>
          </div>
        </div>
      </Card>
    </div>
  );
}
