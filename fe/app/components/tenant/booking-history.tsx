import { Card } from '@/app/components/ui/card';
import { Button } from '@/app/components/ui/button';
import { Badge } from '@/app/components/ui/badge';
import { ImageWithFallback } from '@/app/components/shared/ImageWithFallback';
import { motion } from 'framer-motion';
import { useState } from 'react';
import { ExtendBooking } from './extend-booking';
import { CancelBooking } from './cancel-booking';
import {
  Calendar,
  MapPin,
  Eye,
  Clock,
  Search,
  CheckCircle2,
  XCircle,
  TrendingUp,
  Trash2,
  Loader2
} from 'lucide-react';
import { api } from '@/app/services/api';
import { useEffect } from 'react';

interface Booking {
  id: string;
  roomName: string;
  roomImage: string;
  location: string;
  status: string;
  moveInDate: string;
  moveOutDate: string;
  monthlyRent: number;
  totalPaid: number;
  duration: string;
  rawStatus: string;
}

interface KamarData {
  nomor_kamar: string;
  tipe_kamar: string;
  image_url: string;
  floor: number;
  harga_per_bulan: number;
}


interface BookingHistoryProps {
  onViewRoom: (roomId: string) => void;
}

export function BookingHistory({ onViewRoom }: BookingHistoryProps) {
  const [bookings, setBookings] = useState<Booking[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [extendModalOpen, setExtendModalOpen] = useState(false);
  const [cancelModalOpen, setCancelModalOpen] = useState(false);
  const [selectedBooking, setSelectedBooking] = useState<Booking | null>(null);

  useEffect(() => {
    const fetchBookings = async () => {
      try {
        const data = await api.getMyBookings();
        const mapped = data.map((b: { id: number; kamar: KamarData; tanggal_mulai: string; durasi_sewa: number; status_bayar: string; total_bayar: number }) => {
          const moveIn = new Date(b.tanggal_mulai);
          const moveOut = new Date(moveIn);
          moveOut.setMonth(moveOut.getMonth() + b.durasi_sewa);

          return {
            id: b.id.toString(),
            roomName: b.kamar.nomor_kamar + " - " + b.kamar.tipe_kamar,
            roomImage: b.kamar.image_url.startsWith('http') ? b.kamar.image_url : `http://localhost:8080${b.kamar.image_url}`,
            location: `Floor ${b.kamar.floor}`,
            status: b.status_bayar === 'Confirmed' ? 'Confirmed' : (b.status_bayar === 'Pending' ? 'Pending' : 'Completed'),
            moveInDate: b.tanggal_mulai,
            moveOutDate: moveOut.toISOString().split('T')[0],
            monthlyRent: b.kamar.harga_per_bulan,
            totalPaid: b.total_bayar,
            duration: `${b.durasi_sewa} Months`,
            rawStatus: b.status_bayar
          };
        });
        setBookings(mapped);
      } catch (err) {
        console.error("Failed to fetch bookings", err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchBookings();
  }, []);

  const handleExtendClick = (booking: Booking) => {
    setSelectedBooking(booking);
    setExtendModalOpen(true);
  };

  const handleCancelClick = (booking: Booking) => {
    setSelectedBooking(booking);
    setCancelModalOpen(true);
  };

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'Confirmed': return <CheckCircle2 className="w-5 h-5" />;
      case 'Pending': return <Clock className="w-5 h-5" />;
      case 'Completed': return <CheckCircle2 className="w-5 h-5" />;
      case 'Cancelled': return <XCircle className="w-5 h-5" />;
      default: return null;
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'Confirmed': return 'bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-400 border-emerald-200 dark:border-emerald-800';
      case 'Pending': return 'bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-400 border-amber-200 dark:border-amber-800';
      case 'Completed': return 'bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-400 border-blue-200 dark:border-blue-800';
      case 'Cancelled': return 'bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-400 border-red-200 dark:border-red-800';
      default: return 'bg-slate-100 dark:bg-slate-800 text-slate-700 dark:text-slate-400 border-slate-200 dark:border-slate-700';
    }
  };

  return (
    <div className="bg-white dark:bg-slate-950 text-slate-900 dark:text-slate-50 transition-colors overflow-x-hidden min-h-screen py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Header Section */}
        <motion.div
          initial={{ opacity: 0, y: -30 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6 }}
          className="mb-12"
        >
          <div className="flex items-center gap-3 mb-4">
            <motion.div
              className="w-12 h-12 bg-gradient-to-br from-stone-700 to-stone-900 rounded-lg flex items-center justify-center shadow-lg"
              whileHover={{ scale: 1.1, rotate: 5 }}
            >
              <Calendar className="w-6 h-6 text-white" />
            </motion.div>
            <div>
              <h1 className="text-4xl font-bold text-slate-900 dark:text-white">My Bookings</h1>
              <p className="text-slate-600 dark:text-slate-400 mt-1">Track and manage all your rental reservations</p>
            </div>
          </div>
        </motion.div>

        {/* Summary Cards - 2x2 on Mobile */}
        <motion.div
          className="grid grid-cols-2 lg:grid-cols-4 gap-3 md:gap-6 mb-12"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ staggerChildren: 0.1, delayChildren: 0.2 }}
        >
          {[
            { label: 'Number Room', value: bookings.length.toString(), icon: '🏠', color: 'from-stone-700 to-stone-900', bgColor: 'from-stone-50 to-stone-100 dark:from-slate-900 dark:to-slate-900' },
            { label: 'Bill Completed', value: bookings.filter(b => b.rawStatus === 'Confirmed').length.toString(), icon: '✓', color: 'from-emerald-500 to-emerald-600', bgColor: 'from-emerald-50 to-emerald-100 dark:from-slate-900 dark:to-slate-900' },
            { label: 'Pending Bills', value: bookings.filter(b => b.rawStatus === 'Pending').length.toString(), icon: '⏳', color: 'from-amber-500 to-amber-600', bgColor: 'from-amber-50 to-amber-100 dark:from-slate-900 dark:to-slate-900' },
            { label: 'Total', value: `Rp ${bookings.reduce((sum, b) => sum + b.totalPaid, 0).toLocaleString()}`, icon: '💰', color: 'from-blue-500 to-blue-600', bgColor: 'from-blue-50 to-blue-100 dark:from-slate-900 dark:to-slate-900' },
          ].map((stat, index) => (
            <motion.div
              key={stat.label}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: index * 0.1 }}
              whileHover={{ y: -5 }}
            >
              <Card className={`p-4 md:p-6 bg-gradient-to-br ${stat.bgColor} border-0 dark:border dark:border-slate-800 shadow-lg hover:shadow-xl transition-all duration-300 h-full`}>
                <div className={`w-10 h-10 md:w-14 md:h-14 bg-gradient-to-br ${stat.color} rounded-xl flex items-center justify-center mb-3 md:mb-4 shadow-lg`}>
                  <span className="text-xl md:text-2xl">{stat.icon}</span>
                </div>
                <p className="text-[10px] md:text-sm text-slate-600 dark:text-slate-400 font-semibold mb-1 md:mb-2 uppercase tracking-wide">{stat.label}</p>
                <p className="text-lg md:text-3xl font-bold text-slate-900 dark:text-white line-clamp-1">{stat.value}</p>
              </Card>
            </motion.div>
          ))}
        </motion.div>

        {/* Bookings List */}
        <motion.div
          className="space-y-6"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ staggerChildren: 0.15, delayChildren: 0.4 }}
        >
          {isLoading ? (
            <div className="flex flex-col items-center justify-center py-20 gap-4">
              <Loader2 className="w-10 h-10 text-stone-600 animate-spin" />
              <p className="text-slate-500 font-medium italic">Sinking your luxury reservations...</p>
            </div>
          ) : bookings.map((booking, index) => (
            <motion.div
              key={booking.id}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: index * 0.15 }}
              whileHover={{ y: -5 }}
            >
              <Card className="overflow-hidden hover:shadow-2xl transition-all duration-300 border-0 dark:border dark:border-slate-800 shadow-lg bg-white dark:bg-slate-900">
                <div className="flex flex-col md:flex-row">
                  {/* Room Image with Overlay */}
                  <motion.div
                    className="md:w-80 h-56 md:h-auto flex-shrink-0 relative overflow-hidden bg-gradient-to-br from-slate-200 to-slate-300 dark:from-slate-800 dark:to-slate-900"
                    whileHover={{ scale: 1.05 }}
                  >
                    <ImageWithFallback
                      src={booking.roomImage}
                      alt={booking.roomName}
                      className="w-full h-full object-cover"
                    />
                    <div className="absolute inset-0 bg-gradient-to-t from-black/40 to-transparent opacity-0 hover:opacity-100 transition-opacity duration-300" />
                  </motion.div>

                  {/* Booking Details */}
                  <div className="flex-1 p-8">
                    <div className="flex items-start justify-between mb-6">
                      <div>
                        <h3 className="text-2xl font-bold text-slate-900 dark:text-white mb-2">
                          {booking.roomName}
                        </h3>
                        <div className="flex items-center gap-2 text-slate-600 dark:text-slate-400">
                          <MapPin className="w-4 h-4 text-stone-900 dark:text-stone-400" />
                          <span className="font-medium">{booking.location}</span>
                        </div>
                      </div>
                      <motion.div whileHover={{ scale: 1.1 }} whileTap={{ scale: 0.95 }}>
                        <Badge className={`${getStatusColor(booking.status)} border-2 flex items-center gap-2 px-4 py-2 font-semibold text-sm shadow-md hover:shadow-lg transition-all`}>
                          {getStatusIcon(booking.status)}
                          {booking.status}
                        </Badge>
                      </motion.div>
                    </div>

                    <div className="grid grid-cols-2 lg:grid-cols-4 gap-3 md:gap-4 mb-6">
                      {/* Check-in Card */}
                      <motion.div
                        className="p-3 md:p-4 bg-gradient-to-br from-slate-50 to-slate-100 dark:from-slate-800/50 dark:to-slate-800/80 rounded-xl border border-slate-200/50 dark:border-slate-700/50 hover:shadow-md transition-all"
                        whileHover={{ y: -2 }}
                      >
                        <p className="text-[10px] md:text-xs text-slate-600 dark:text-slate-400 font-semibold mb-1 md:mb-2 uppercase tracking-wide">Check-in</p>
                        <div className="flex items-center gap-1.5 md:gap-2 text-slate-900 dark:text-white">
                          <Calendar className="w-4 h-4 md:w-5 md:h-5 text-stone-900 dark:text-slate-400" />
                          <span className="font-semibold text-xs md:text-base">{new Date(booking.moveInDate).toLocaleDateString('en-GB')}</span>
                        </div>
                      </motion.div>

                      {/* Check-out Card */}
                      <motion.div
                        className="p-3 md:p-4 bg-gradient-to-br from-slate-50 to-slate-100 dark:from-slate-800/50 dark:to-slate-800/80 rounded-xl border border-slate-200/50 dark:border-slate-700/50 hover:shadow-md transition-all"
                        whileHover={{ y: -2 }}
                      >
                        <p className="text-[10px] md:text-xs text-slate-600 dark:text-slate-400 font-semibold mb-1 md:mb-2 uppercase tracking-wide">Check-out</p>
                        <div className="flex items-center gap-1.5 md:gap-2 text-slate-900 dark:text-white">
                          <Calendar className="w-4 h-4 md:w-5 md:h-5 text-stone-900 dark:text-slate-400" />
                          <span className="font-semibold text-xs md:text-base">{new Date(booking.moveOutDate).toLocaleDateString('en-GB')}</span>
                        </div>
                      </motion.div>

                      {/* Duration Card */}
                      <motion.div
                        className="p-3 md:p-4 bg-gradient-to-br from-purple-50 to-purple-100 dark:from-purple-900/20 dark:to-purple-900/30 rounded-xl border border-purple-200/50 dark:border-purple-800/50 hover:shadow-md transition-all"
                        whileHover={{ y: -2 }}
                      >
                        <p className="text-[10px] md:text-xs text-purple-700 dark:text-purple-400 font-semibold mb-1 md:mb-2 uppercase tracking-wide">Duration</p>
                        <div className="flex items-center gap-1.5 md:gap-2">
                          <Clock className="w-4 h-4 md:w-5 md:h-5 text-purple-700 dark:text-purple-400" />
                          <span className="font-semibold text-xs md:text-base text-purple-900 dark:text-purple-200">{booking.duration}</span>
                        </div>
                      </motion.div>

                      {/* Monthly Card */}
                      <motion.div
                        className="p-3 md:p-4 bg-gradient-to-br from-emerald-50 to-emerald-100 dark:from-emerald-900/20 dark:to-emerald-900/30 rounded-xl border border-emerald-200/50 dark:border-emerald-800/50 hover:shadow-md transition-all"
                        whileHover={{ y: -2 }}
                      >
                        <p className="text-[10px] md:text-xs text-emerald-700 dark:text-emerald-400 font-semibold mb-1 md:mb-2 uppercase tracking-wide">Monthly</p>
                        <div className="flex items-center gap-1">
                          <span className="font-bold text-emerald-900 dark:text-emerald-200 text-[10px] md:text-sm">Rp</span>
                          <span className="font-bold text-emerald-900 dark:text-emerald-200 text-sm md:text-lg">{booking.monthlyRent.toLocaleString()}</span>
                        </div>
                      </motion.div>
                    </div>

                    <div className="flex items-center justify-between pt-6 border-t border-slate-200 dark:border-slate-800">
                      <motion.div whileHover={{ scale: 1.05 }}>
                        <div>
                          <p className="text-sm text-slate-600 dark:text-slate-400 font-medium mb-1">Total Amount Paid</p>
                          <p className="text-4xl font-bold bg-gradient-to-r from-stone-700 to-stone-900 dark:from-white dark:to-slate-400 bg-clip-text text-transparent">
                            Rp {booking.totalPaid.toLocaleString()}
                          </p>
                        </div>
                      </motion.div>

                      <div className="flex gap-3 flex-wrap">
                        <motion.div whileHover={{ scale: 1.05 }} whileTap={{ scale: 0.95 }}>
                          <Button
                            variant="outline"
                            onClick={() => onViewRoom(booking.id)}
                            className="border-slate-300 dark:border-slate-700 hover:bg-slate-100 dark:hover:bg-slate-800 font-semibold shadow-md hover:shadow-lg transition-all dark:text-white"
                          >
                            <Eye className="w-4 h-4 mr-2" />
                            View Details
                          </Button>
                        </motion.div>
                        {booking.status === 'Confirmed' && (
                          <>
                            <motion.div whileHover={{ scale: 1.05 }} whileTap={{ scale: 0.95 }}>
                              <Button
                                onClick={() => handleExtendClick(booking)}
                                className="bg-gradient-to-r from-stone-700 to-stone-900 dark:from-stone-600 dark:to-stone-800 hover:from-stone-600 hover:to-stone-800 text-white font-semibold shadow-lg hover:shadow-xl transition-all"
                              >
                                <TrendingUp className="w-4 h-4 mr-2" />
                                Extend Booking
                              </Button>
                            </motion.div>
                            <motion.div whileHover={{ scale: 1.05 }} whileTap={{ scale: 0.95 }}>
                              <Button
                                onClick={() => handleCancelClick(booking)}
                                variant="outline"
                                className="border-red-300 dark:border-red-900/50 text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 hover:text-red-700 font-semibold shadow-md hover:shadow-lg transition-all"
                              >
                                <Trash2 className="w-4 h-4 mr-2" />
                                Cancel
                              </Button>
                            </motion.div>
                          </>
                        )}
                      </div>
                    </div>
                  </div>
                </div>
              </Card>
            </motion.div>
          ))}
        </motion.div>

        {/* Empty State */}
        {!isLoading && bookings.length === 0 && (
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6 }}
          >
            <Card className="p-16 text-center bg-gradient-to-br from-slate-50 to-stone-50 dark:from-slate-900 dark:to-slate-950 border-0 dark:border dark:border-slate-800 shadow-lg">
              <motion.div
                className="w-20 h-20 bg-gradient-to-br from-stone-700 to-stone-900 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg"
                whileHover={{ scale: 1.1, rotate: 10 }}
              >
                <Calendar className="w-10 h-10 text-white" />
              </motion.div>
              <h3 className="text-3xl font-bold text-slate-900 dark:text-white mb-3">No bookings yet</h3>
              <p className="text-slate-600 dark:text-slate-400 mb-8 text-lg max-w-md mx-auto">
                Start your journey to premium living. Explore our collection of luxury boarding houses and apartments.
              </p>
              <motion.div whileHover={{ scale: 1.05 }} whileTap={{ scale: 0.95 }}>
                <Button className="bg-gradient-to-r from-stone-700 to-stone-900 dark:from-stone-600 dark:to-stone-800 text-white font-semibold px-8 py-2.5 shadow-lg hover:shadow-xl transition-all">
                  <Search className="w-4 h-4 mr-2" />
                  Browse Rooms
                </Button>
              </motion.div>
            </Card>
          </motion.div>
        )}
      </div>

      {/* Modals remain same, they should handle dark mode internally via their components */}
      {selectedBooking && (
        <ExtendBooking
          isOpen={extendModalOpen}
          onClose={() => setExtendModalOpen(false)}
          bookingData={{
            id: selectedBooking.id,
            roomName: selectedBooking.roomName,
            currentEndDate: selectedBooking.moveOutDate,
            pricePerMonth: selectedBooking.monthlyRent,
            image: selectedBooking.roomImage,
          }}
        />
      )}

      {selectedBooking && (
        <CancelBooking
          isOpen={cancelModalOpen}
          onClose={() => setCancelModalOpen(false)}
          bookingData={{
            id: selectedBooking.id,
            roomName: selectedBooking.roomName,
            moveOutDate: selectedBooking.moveOutDate,
            monthlyRent: selectedBooking.monthlyRent,
            totalPaid: selectedBooking.totalPaid,
            image: selectedBooking.roomImage,
            duration: selectedBooking.duration,
            status: selectedBooking.status as 'Confirmed' | 'Pending' | 'Completed' | 'Cancelled',
          }}
        />
      )}
    </div>
  );
}