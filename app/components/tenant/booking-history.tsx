import { Card } from '@/app/components/ui/card';
import { Button } from '@/app/components/ui/button';
import { Badge } from '@/app/components/ui/badge';
import { ImageWithFallback } from '@/app/components/figma/ImageWithFallback';
import { 
  Calendar, 
  MapPin, 
  DollarSign,
  Eye,
  Clock,
  Search,
  CheckCircle2,
  XCircle
} from 'lucide-react';

interface Booking {
  id: string;
  roomName: string;
  roomImage: string;
  location: string;
  status: 'Confirmed' | 'Pending' | 'Completed' | 'Cancelled';
  moveInDate: string;
  moveOutDate: string;
  monthlyRent: number;
  totalPaid: number;
  duration: string;
}

const mockBookings: Booking[] = [
  {
    id: '1',
    roomName: 'Deluxe Suite #12',
    roomImage: 'https://images.unsplash.com/photo-1668512624222-2e375314be39?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxsdXh1cnklMjBib2FyZGluZyUyMGhvdXNlJTIwYmVkcm9vbXxlbnwxfHx8fDE3Njg1NzcyMzF8MA&ixlib=rb-4.1.0&q=80&w=400',
    location: 'Downtown District',
    status: 'Confirmed',
    moveInDate: '2026-02-01',
    moveOutDate: '2026-08-01',
    monthlyRent: 1200,
    totalPaid: 2400,
    duration: '6 months'
  },
  {
    id: '2',
    roomName: 'Modern Studio #24',
    roomImage: 'https://images.unsplash.com/photo-1662454419736-de132ff75638?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxtb2Rlcm4lMjBhcGFydG1lbnQlMjBiZWRyb29tJTIwaW50ZXJpb3J8ZW58MXx8fHwxNzY4NTc3MjMxfDA&ixlib=rb-4.1.0&q=80&w=400',
    location: 'University Area',
    status: 'Pending',
    moveInDate: '2026-03-01',
    moveOutDate: '2026-06-01',
    monthlyRent: 800,
    totalPaid: 1600,
    duration: '3 months'
  },
  {
    id: '3',
    roomName: 'Premium Apartment #8',
    roomImage: 'https://images.unsplash.com/photo-1507138451611-3001135909fa?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxjb3p5JTIwc3R1ZGlvJTIwYXBhcnRtZW50fGVufDF8fHx8MTc2ODUwNTcxM3ww&ixlib=rb-4.1.0&q=80&w=400',
    location: 'Business District',
    status: 'Completed',
    moveInDate: '2025-08-01',
    moveOutDate: '2026-01-01',
    monthlyRent: 1500,
    totalPaid: 9000,
    duration: '6 months'
  },
];

interface BookingHistoryProps {
  onViewRoom: (roomId: string) => void;
}

export function BookingHistory({ onViewRoom }: BookingHistoryProps) {
  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'Confirmed':
        return <CheckCircle2 className="w-5 h-5" />;
      case 'Pending':
        return <Clock className="w-5 h-5" />;
      case 'Completed':
        return <CheckCircle2 className="w-5 h-5" />;
      case 'Cancelled':
        return <XCircle className="w-5 h-5" />;
      default:
        return null;
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'Confirmed':
        return 'bg-emerald-100 text-emerald-700 border-emerald-200';
      case 'Pending':
        return 'bg-amber-100 text-amber-700 border-amber-200';
      case 'Completed':
        return 'bg-blue-100 text-blue-700 border-blue-200';
      case 'Cancelled':
        return 'bg-red-100 text-red-700 border-red-200';
      default:
        return 'bg-slate-100 text-slate-700 border-slate-200';
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 via-slate-100 to-stone-50 py-12">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Header Section */}
        <div className="mb-12">
          <div className="inline-block mb-4">
            <Badge className="bg-stone-900/10 text-stone-900 font-semibold">My Bookings</Badge>
          </div>
          <h1 className="text-5xl font-bold text-slate-900 mb-3">Booking History</h1>
          <p className="text-lg text-slate-600">Manage and view all your rental reservations</p>
        </div>

        {/* Summary Cards */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-12">
          {[
            { label: 'Total Bookings', value: '3', icon: '📅', color: 'from-stone-700 to-stone-900' },
            { label: 'Active', value: '1', icon: '✓', color: 'from-emerald-500 to-emerald-600' },
            { label: 'Pending', value: '1', icon: '⏳', color: 'from-amber-500 to-amber-600' },
            { label: 'Total Spent', value: '$13,000', icon: '💰', color: 'from-blue-500 to-blue-600' },
          ].map((stat, index) => (
            <div key={stat.label}>
              <Card className="p-6 bg-white/80 backdrop-blur-sm border-slate-200/60 hover:shadow-lg transition-all duration-300">
                <div className={`w-14 h-14 bg-gradient-to-br ${stat.color} rounded-xl flex items-center justify-center mb-4 shadow-lg`}>
                  <span className="text-2xl">{stat.icon}</span>
                </div>
                <p className="text-sm text-slate-600 font-medium mb-2">{stat.label}</p>
                <p className="text-3xl font-bold text-slate-900">{stat.value}</p>
              </Card>
            </div>
          ))}
        </div>

        {/* Bookings List */}
        <div className="space-y-6">
          {mockBookings.map((booking, index) => (
            <div key={booking.id}>
              <Card className="overflow-hidden hover:shadow-xl transition-all duration-300 border-slate-200/60 bg-white/90 backdrop-blur-sm">
                <div className="flex flex-col md:flex-row">
                  {/* Room Image with Overlay */}
                  <div className="md:w-72 h-56 md:h-auto flex-shrink-0 relative overflow-hidden">
                    <ImageWithFallback
                      src={booking.roomImage}
                      alt={booking.roomName}
                      className="w-full h-full object-cover hover:scale-105 transition-transform duration-300"
                    />
                    <div className="absolute inset-0 bg-gradient-to-t from-slate-900/20 to-transparent opacity-0 hover:opacity-100 transition-opacity duration-300" />
                  </div>

                  {/* Booking Details */}
                  <div className="flex-1 p-7">
                    <div className="flex items-start justify-between mb-5">
                      <div>
                        <h3 className="text-2xl font-bold text-slate-900 mb-2">
                          {booking.roomName}
                        </h3>
                        <div className="flex items-center gap-2 text-slate-600">
                          <MapPin className="w-4 h-4 text-stone-700" />
                          <span className="font-medium">{booking.location}</span>
                        </div>
                      </div>
                      <Badge className={`${getStatusColor(booking.status)} border-2 flex items-center gap-2 px-4 py-1.5 font-semibold text-sm`}>
                        {getStatusIcon(booking.status)}
                        {booking.status}
                      </Badge>
                    </div>

                    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-5">
                      <div className="p-4 bg-gradient-to-br from-slate-50 to-slate-100 rounded-xl border border-slate-200/50">
                        <p className="text-xs text-slate-600 font-semibold mb-2 uppercase tracking-wide">Check-in</p>
                        <div className="flex items-center gap-2">
                          <Calendar className="w-5 h-5 text-stone-700" />
                          <span className="font-semibold text-slate-900">
                            {new Date(booking.moveInDate).toLocaleDateString('en-US', {
                              month: 'short',
                              day: 'numeric',
                              year: 'numeric'
                            })}
                          </span>
                        </div>
                      </div>

                      <div className="p-4 bg-gradient-to-br from-slate-50 to-slate-100 rounded-xl border border-slate-200/50">
                        <p className="text-xs text-slate-600 font-semibold mb-2 uppercase tracking-wide">Check-out</p>
                        <div className="flex items-center gap-2">
                          <Calendar className="w-5 h-5 text-stone-700" />
                          <span className="font-semibold text-slate-900">
                            {new Date(booking.moveOutDate).toLocaleDateString('en-US', {
                              month: 'short',
                              day: 'numeric',
                              year: 'numeric'
                            })}
                          </span>
                        </div>
                      </div>

                      <div className="p-4 bg-gradient-to-br from-slate-50 to-slate-100 rounded-xl border border-slate-200/50">
                        <p className="text-xs text-slate-600 font-semibold mb-2 uppercase tracking-wide">Duration</p>
                        <div className="flex items-center gap-2">
                          <Clock className="w-5 h-5 text-stone-700" />
                          <span className="font-semibold text-slate-900">{booking.duration}</span>
                        </div>
                      </div>

                      <div className="p-4 bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl border border-blue-200/50">
                        <p className="text-xs text-blue-700 font-semibold mb-2 uppercase tracking-wide">Monthly</p>
                        <div className="flex items-center gap-1">
                          <DollarSign className="w-5 h-5 text-blue-700" />
                          <span className="font-bold text-blue-900 text-lg">
                            {booking.monthlyRent.toLocaleString()}
                          </span>
                        </div>
                      </div>
                    </div>

                    <div className="flex items-center justify-between pt-5 border-t border-slate-200">
                      <div>
                        <p className="text-sm text-slate-600 font-medium mb-1">Amount Paid</p>
                        <p className="text-3xl font-bold bg-gradient-to-r from-stone-700 to-stone-900 bg-clip-text text-transparent">
                          ${booking.totalPaid.toLocaleString()}
                        </p>
                      </div>

                      <div className="flex gap-3">
                        <Button
                          variant="outline"
                          onClick={() => onViewRoom(booking.id)}
                          className="border-slate-300 hover:bg-slate-50 font-medium"
                        >
                          <Eye className="w-4 h-4 mr-2" />
                          View Details
                        </Button>
                        {booking.status === 'Confirmed' && (
                          <Button className="bg-stone-900 hover:bg-stone-800 text-white font-medium shadow-md hover:shadow-lg transition-all">
                            <Calendar className="w-4 h-4 mr-2" />
                            Extend Booking
                          </Button>
                        )}
                      </div>
                    </div>
                  </div>
                </div>
              </Card>
            </div>
          ))}
        </div>

        {/* Empty State (if no bookings) */}
        {mockBookings.length === 0 && (
          <Card className="p-16 text-center bg-gradient-to-br from-slate-50 to-stone-50 border-slate-200/60">
            <div className="w-20 h-20 bg-gradient-to-br from-stone-700 to-stone-900 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg">
              <Calendar className="w-10 h-10 text-white" />
            </div>
            <h3 className="text-3xl font-bold text-slate-900 mb-3">No bookings yet</h3>
            <p className="text-slate-600 mb-8 text-lg max-w-md mx-auto">
              Start your journey to premium living. Explore our collection of luxury boarding houses and apartments.
            </p>
            <Button className="bg-stone-900 hover:bg-stone-800 text-white font-semibold px-8 py-2 shadow-lg hover:shadow-xl transition-all">
              <Search className="w-4 h-4 mr-2" />
              Browse Rooms
            </Button>
          </Card>
        )}
      </div>
    </div>
  );
}