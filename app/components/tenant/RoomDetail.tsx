import { useState } from 'react';
import { Card } from '@/app/components/ui/card';
import { Button } from '@/app/components/ui/button';
import { Badge } from '@/app/components/ui/badge';
import { ImageWithFallback } from '@/app/components/figma/ImageWithFallback';
import { 
  ArrowLeft, 
  MapPin, 
  Star, 
  Wifi, 
  Wind, 
  Tv, 
  Coffee,
  Bed,
  Bath,
  Ruler,
  ChevronLeft,
  ChevronRight,
  Check,
  Calendar,
  Utensils,
  Monitor,
  Briefcase
} from 'lucide-react';

interface RoomDetailProps {
  roomId: string;
  onBookNow: (roomId: string) => void;
  onBack: () => void;
}

// Data Kamar
const roomDetails: { [key: string]: any } = {
  '1': {
    name: 'Premium Suite 201',
    type: 'Luxury',
    price: 1200,
    location: 'South Jakarta',
    rating: 4.9,
    reviews: 98,
    description: 'Luxurious premium suite with full kitchen and spacious balcony. Ideal for those who appreciate the finer things.',
    bedrooms: 1,
    bathrooms: 1,
    size: '45m²',
    facilities: [
      { name: 'High-Speed WiFi', icon: Wifi },
      { name: 'Full Kitchen', icon: Utensils },
      { name: 'Air conditioning', icon: Wind },
      { name: 'Smart TV', icon: Monitor },
      { name: 'King Size Bed', icon: Bed },
      { name: 'Work Desk', icon: Briefcase },
    ],
    features: [
      'Air conditioning',
      'Private bathroom',
      'Hot water',
      'Work desk',
      'High-speed WiFi',
      'Smart TV with cable',
      'Fully equipped kitchen',
      'Balcony with city view',
      'Premium bedding',
      'Daily housekeeping',
    ],
    images: [
      'https://images.unsplash.com/photo-1668512624222-2e375314be39?q=80&w=1080',
      'https://images.unsplash.com/photo-1662454419736-de132ff75638?q=80&w=1080',
      'https://images.unsplash.com/photo-1507138451611-3001135909fa?q=80&w=1080',
    ],
  },
};

const facilityIcons: { [key: string]: any } = {
  'High-Speed WiFi': Wifi,
  'Air conditioning': Wind,
  'Smart TV': Monitor,
  'King Size Bed': Bed,
  'Work Desk': Briefcase,
  'Full Kitchen': Utensils,
};

export function RoomDetail({ roomId, onBookNow, onBack }: RoomDetailProps) {
  const [currentImageIndex, setCurrentImageIndex] = useState(0);
  const room = roomDetails[roomId] || roomDetails['1'];

  const nextImage = () => {
    setCurrentImageIndex((prev) => (prev + 1) % room.images.length);
  };

  const prevImage = () => {
    setCurrentImageIndex((prev) => (prev - 1 + room.images.length) % room.images.length);
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-stone-100 py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        
        {/* Top Header - Back Button & Title */}
        <div className="flex items-center gap-4 mb-8">
          <Button variant="ghost" onClick={onBack} className="rounded-full h-10 w-10 p-0">
            <ArrowLeft className="w-5 h-5" />
          </Button>
          <div>
            <h1 className="text-2xl font-bold text-slate-900">{room.name}</h1>
            <div className="flex items-center gap-1 text-slate-500 text-sm">
              <MapPin className="w-4 h-4" />
              <span>{room.location}</span>
            </div>
          </div>
        </div>

        {/* Content Grid */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          
          {/* Left Column: Image Gallery & Details */}
          <div className="lg:col-span-2 space-y-6">
            
            {/* Main Image Gallery */}
            <div className="space-y-4">
              <Card className="overflow-hidden">
                <div className="relative h-[450px] bg-slate-900">
                  <ImageWithFallback
                    src={room.images[currentImageIndex]}
                    alt={room.name}
                    className="w-full h-full object-cover"
                  />
                  
                  {/* Navigation Arrows */}
                  {room.images.length > 1 && (
                    <>
                      <button
                        onClick={prevImage}
                        className="absolute left-4 top-1/2 -translate-y-1/2 w-12 h-12 bg-white/90 rounded-full shadow-lg flex items-center justify-center hover:bg-white transition-all"
                      >
                        <ChevronLeft className="w-6 h-6 text-slate-900" />
                      </button>
                      <button
                        onClick={nextImage}
                        className="absolute right-4 top-1/2 -translate-y-1/2 w-12 h-12 bg-white/90 rounded-full shadow-lg flex items-center justify-center hover:bg-white transition-all"
                      >
                        <ChevronRight className="w-6 h-6 text-slate-900" />
                      </button>
                    </>
                  )}
                </div>
              </Card>

              {/* Thumbnail Gallery */}
              <div className="grid grid-cols-3 gap-3">
                {room.images.map((img: string, idx: number) => (
                  <div 
                    key={idx}
                    onClick={() => setCurrentImageIndex(idx)}
                    className={`h-20 rounded-lg overflow-hidden cursor-pointer transition-all border-2 ${
                      idx === currentImageIndex ? 'border-amber-500' : 'border-slate-200'
                    }`}
                  >
                    <img src={img} alt="thumbnail" className="w-full h-full object-cover" />
                  </div>
                ))}
              </div>
            </div>

            {/* Room Info Card */}
            <Card className="p-6">
              <div className="mb-4">
                <Badge className="bg-stone-900 text-white">{room.type}</Badge>
              </div>
              
              <div className="mb-6 pb-6 border-b border-slate-200">
                <div className="flex items-center gap-4 mb-4">
                  <div className="flex items-center gap-2">
                    <Star className="w-5 h-5 fill-amber-400 text-amber-400" />
                    <span className="font-bold">{room.rating}</span>
                    <span className="text-slate-600">({room.reviews} reviews)</span>
                  </div>
                </div>

                <div className="flex gap-6 text-sm">
                  <div className="flex items-center gap-2">
                    <Bed className="w-5 h-5 text-slate-600" />
                    <span>{room.bedrooms} Bed</span>
                  </div>
                  <div className="flex items-center gap-2">
                    <Bath className="w-5 h-5 text-slate-600" />
                    <span>{room.bathrooms} Bath</span>
                  </div>
                  <div className="flex items-center gap-2">
                    <Ruler className="w-5 h-5 text-slate-600" />
                    <span>{room.size}</span>
                  </div>
                </div>
              </div>

              <h3 className="font-bold text-lg mb-3">Facilities</h3>
              <div className="grid grid-cols-2 gap-3">
                {room.facilities.map((facility: any, idx: number) => {
                  const Icon = facilityIcons[facility.name] || facility.icon;
                  return (
                    <div key={idx} className="flex items-center gap-2 p-2 bg-slate-50 rounded-lg">
                      {Icon && <Icon className="w-4 h-4 text-slate-600" />}
                      <span className="text-sm text-slate-700">{facility.name}</span>
                    </div>
                  );
                })}
              </div>
            </Card>

            {/* Description */}
            <Card className="p-6">
              <h2 className="text-2xl font-bold text-slate-900 mb-4">About This Room</h2>
              <p className="text-slate-600 leading-relaxed">{room.description}</p>
            </Card>

            {/* Features */}
            <Card className="p-6">
              <h2 className="text-2xl font-bold text-slate-900 mb-4">Features & Amenities</h2>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-3">
                {room.features.map((feature: string, idx: number) => (
                  <div key={idx} className="flex items-start gap-2">
                    <Check className="w-5 h-5 text-emerald-600 flex-shrink-0 mt-0.5" />
                    <span className="text-slate-600">{feature}</span>
                  </div>
                ))}
              </div>
            </Card>
          </div>

          {/* Right Column: Booking Sidebar */}
          <div className="lg:col-span-1">
            <Card className="p-6 sticky top-8 shadow-xl">
              <div className="mb-6">
                <div className="text-4xl font-bold text-slate-900 mb-2">
                  ${room.price}
                  <span className="text-lg text-slate-600">/month</span>
                </div>
              </div>

              {/* Calendar */}
              <div className="mb-6">
                <div className="flex items-center gap-2 mb-4">
                  <Calendar className="w-5 h-5 text-amber-500" />
                  <h3 className="font-bold text-slate-900">Check Availability</h3>
                </div>
                
                <div className="bg-slate-50 border border-slate-200 rounded-xl p-4">
                  <div className="flex justify-between items-center mb-4">
                    <button className="p-1 hover:bg-slate-200 rounded-lg">
                      <ChevronLeft className="w-4 h-4" />
                    </button>
                    <span className="font-bold text-sm">January 2026</span>
                    <button className="p-1 hover:bg-slate-200 rounded-lg">
                      <ChevronRight className="w-4 h-4" />
                    </button>
                  </div>
                  
                  <div className="grid grid-cols-7 gap-1 text-center text-xs mb-2 text-slate-500">
                    <span>Su</span>
                    <span>Mo</span>
                    <span>Tu</span>
                    <span>We</span>
                    <span>Th</span>
                    <span>Fr</span>
                    <span>Sa</span>
                  </div>
                  
                  <div className="grid grid-cols-7 gap-1">
                    {[...Array(31)].map((_, i) => (
                      <div 
                        key={i} 
                        className={`py-2 text-xs rounded cursor-pointer transition-colors ${
                          i === 15 ? 'bg-amber-500 text-white font-bold' : 'hover:bg-slate-200'
                        }`}
                      >
                        {i + 1}
                      </div>
                    ))}
                  </div>
                </div>
              </div>

              <Button 
                onClick={() => onBookNow(roomId)}
                className="w-full bg-amber-500 hover:bg-amber-600 h-12 text-lg font-bold rounded-xl mb-6"
              >
                Book Now
              </Button>

              <div className="space-y-3 border-t border-slate-200 pt-4">
                <div className="flex items-center justify-between text-sm">
                  <span className="text-slate-600">Monthly rent</span>
                  <span className="font-medium">${room.price}</span>
                </div>
                <div className="flex items-center justify-between text-sm">
                  <span className="text-slate-600">Security deposit</span>
                  <span className="font-medium">${room.price}</span>
                </div>
                <div className="flex items-center justify-between text-sm pt-3 border-t border-slate-200 font-bold">
                  <span>Total due</span>
                  <span className="text-lg">${room.price * 2}</span>
                </div>
              </div>

              <div className="mt-4 p-3 bg-green-50 rounded-lg border border-green-200">
                <p className="text-xs text-green-700 text-center">
                  ✓ Secure booking • Free cancellation within 24 hours
                </p>
              </div>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
}