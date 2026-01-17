import { useState } from 'react';
import { Card } from '@/app/components/ui/card';
import { Button } from '@/app/components/ui/button';
import { Input } from '@/app/components/ui/input';
import { Badge } from '@/app/components/ui/badge';
import { ImageWithFallback } from '@/app/components/figma/ImageWithFallback';
import { motion } from 'framer-motion';
import Slider from 'react-slick';
import { 
  Search, 
  MapPin, 
  Wifi, 
  Wind, 
  Tv, 
  Coffee, 
  Star,
  ChevronLeft,
  ChevronRight
} from 'lucide-react';


interface Room {
  id: string;
  name: string;
  type: string;
  price: number;
  image: string;
  location: string;
  rating: number;
  reviews: number;
  facilities: string[];
}

const featuredRooms: Room[] = [
  {
    id: '1',
    name: 'Deluxe Suite',
    type: 'Deluxe',
    price: 1200,
    image: 'https://images.unsplash.com/photo-1668512624222-2e375314be39?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxsdXh1cnklMjBib2FyZGluZyUyMGhvdXNlJTIwYmVkcm9vbXxlbnwxfHx8fDE3Njg1NzcyMzF8MA&ixlib=rb-4.1.0&q=80&w=1080',
    location: 'Downtown District',
    rating: 4.8,
    reviews: 124,
    facilities: ['WiFi', 'AC', 'TV', 'Coffee Maker']
  },
  {
    id: '2',
    name: 'Modern Studio',
    type: 'Standard',
    price: 800,
    image: 'https://images.unsplash.com/photo-1662454419736-de132ff75638?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxtb2Rlcm4lMjBhcGFydG1lbnQlMjBiZWRyb29tJTIwaW50ZXJpb3J8ZW58MXx8fHwxNzY4NTc3MjMxfDA&ixlib=rb-4.1.0&q=80&w=1080',
    location: 'University Area',
    rating: 4.6,
    reviews: 89,
    facilities: ['WiFi', 'AC', 'TV']
  },
  {
    id: '3',
    name: 'Premium Apartment',
    type: 'Premium',
    price: 1500,
    image: 'https://images.unsplash.com/photo-1507138451611-3001135909fa?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxjb3p5JTIwc3R1ZGlvJTIwYXBhcnRtZW50fGVufDF8fHx8MTc2ODUwNTcxM3ww&ixlib=rb-4.1.0&q=80&w=1080',
    location: 'Business District',
    rating: 4.9,
    reviews: 156,
    facilities: ['WiFi', 'AC', 'TV', 'Coffee Maker']
  },
  {
    id: '4',
    name: 'Executive Suite',
    type: 'Executive',
    price: 2000,
    image: 'https://images.unsplash.com/photo-1661258279966-cfeb51c98327?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxwcmVtaXVtJTIwaG90ZWwlMjByb29tfGVufDF8fHx8MTc2ODUzMDkzMXww&ixlib=rb-4.1.0&q=80&w=1080',
    location: 'Luxury Quarter',
    rating: 5.0,
    reviews: 203,
    facilities: ['WiFi', 'AC', 'TV', 'Coffee Maker']
  },
];

const facilityIcons: { [key: string]: any } = {
  WiFi: Wifi,
  AC: Wind,
  TV: Tv,
  'Coffee Maker': Coffee,
};

// Custom arrow components
function NextArrow(props: any) {
  const { onClick } = props;
  return (
    <button
      onClick={onClick}
      className="absolute right-4 top-1/2 -translate-y-1/2 z-10 w-12 h-12 bg-white rounded-full shadow-lg flex items-center justify-center hover:bg-slate-50 transition-all"
    >
      <ChevronRight className="w-6 h-6 text-slate-900" />
    </button>
  );
}

function PrevArrow(props: any) {
  const { onClick } = props;
  return (
    <button
      onClick={onClick}
      className="absolute left-4 top-1/2 -translate-y-1/2 z-10 w-12 h-12 bg-white rounded-full shadow-lg flex items-center justify-center hover:bg-slate-50 transition-all"
    >
      <ChevronLeft className="w-6 h-6 text-slate-900" />
    </button>
  );
}

interface HomepageProps {
  onRoomClick: (roomId: string) => void;
}

export function Homepage({ onRoomClick }: HomepageProps) {
  const [searchLocation, setSearchLocation] = useState('');
  const [priceRange, setPriceRange] = useState([0, 3000]);

  const sliderSettings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 3,
    slidesToScroll: 1,
    autoplay: true,
    autoplaySpeed: 4000,
    nextArrow: <NextArrow />,
    prevArrow: <PrevArrow />,
    responsive: [
      {
        breakpoint: 1024,
        settings: {
          slidesToShow: 2,
        }
      },
      {
        breakpoint: 640,
        settings: {
          slidesToShow: 1,
        }
      }
    ]
  };

  return (
    <div>
      {/* Hero Section */}
      <section className="relative bg-gradient-to-r from-stone-800 via-stone-700 to-stone-600 text-white py-20 overflow-hidden">
        <div className="absolute inset-0 bg-[url('https://images.unsplash.com/photo-1721009714214-e688d8c07506?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w3Nzg4Nzd8MHwxfHNlYXJjaHwxfHxlbGVnYW50JTIwcm9vbSUyMGludGVyaW9yfGVufDF8fHx8MTc2ODU3NzIzM3ww&ixlib=rb-4.1.0&q=80&w=1080')] bg-cover bg-center opacity-20" />
        <div className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6 }}
            className="text-center mb-12"
          >
            <h1 className="text-5xl md:text-6xl font-bold mb-4">
              Find Your Perfect Space
            </h1>
            <p className="text-xl text-stone-200">
              Premium boarding houses tailored to your lifestyle
            </p>
          </motion.div>

          {/* Search Bar */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6, delay: 0.2 }}
          >
            <Card className="bg-white/95 backdrop-blur-xl p-6 shadow-2xl">
              <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
                <div className="md:col-span-2">
                  <label className="text-sm text-slate-600 mb-2 block">Location</label>
                  <div className="relative">
                    <MapPin className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
                    <Input
                      placeholder="Where do you want to stay?"
                      value={searchLocation}
                      onChange={(e) => setSearchLocation(e.target.value)}
                      className="pl-10 border-slate-300"
                    />
                  </div>
                </div>
                <div>
                  <label className="text-sm text-slate-600 mb-2 block">Price Range</label>
                  <select className="w-full px-4 py-2 border border-slate-300 rounded-lg">
                    <option>$0 - $1000</option>
                    <option>$1000 - $2000</option>
                    <option>$2000+</option>
                  </select>
                </div>
                <div>
                  <label className="text-sm text-slate-600 mb-2 block">Room Type</label>
                  <select className="w-full px-4 py-2 border border-slate-300 rounded-lg">
                    <option>All Types</option>
                    <option>Standard</option>
                    <option>Deluxe</option>
                    <option>Premium</option>
                    <option>Executive</option>
                  </select>
                </div>
              </div>
              <Button className="w-full md:w-auto mt-4 bg-stone-900 hover:bg-stone-800">
                <Search className="w-4 h-4 mr-2" />
                Search Rooms
              </Button>
            </Card>
          </motion.div>
        </div>
      </section>

      {/* Featured Rooms Carousel */}
      <section className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6, delay: 0.3 }}
        >
          <div className="text-center mb-12">
            <h2 className="text-4xl font-bold text-slate-900 mb-4">Featured Rooms</h2>
            <p className="text-lg text-slate-600">Handpicked premium accommodations just for you</p>
          </div>

          <div className="relative">
            <Slider {...sliderSettings}>
              {featuredRooms.map((room) => (
                <div key={room.id} className="px-3">
                  <Card className="bg-white shadow-lg hover:shadow-2xl transition-all duration-300 overflow-hidden group cursor-pointer">
                    <div className="relative h-64 overflow-hidden">
                      <ImageWithFallback
                        src={room.image}
                        alt={room.name}
                        className="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500"
                      />
                      <div className="absolute top-4 right-4">
                        <Badge className="bg-stone-900 text-white border-0">
                          {room.type}
                        </Badge>
                      </div>
                      <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/60 to-transparent p-4">
                        <div className="flex items-center gap-2 text-white">
                          <MapPin className="w-4 h-4" />
                          <span className="text-sm">{room.location}</span>
                        </div>
                      </div>
                    </div>

                    <div className="p-5">
                      <div className="flex items-start justify-between mb-3">
                        <h3 className="text-xl font-bold text-slate-900">{room.name}</h3>
                        <div className="flex items-center gap-1">
                          <Star className="w-4 h-4 fill-amber-400 text-amber-400" />
                          <span className="text-sm font-medium">{room.rating}</span>
                          <span className="text-xs text-slate-500">({room.reviews})</span>
                        </div>
                      </div>

                      <div className="flex gap-2 mb-4 flex-wrap">
                        {room.facilities.map((facility) => {
                          const Icon = facilityIcons[facility];
                          return (
                            <div
                              key={facility}
                              className="flex items-center gap-1 px-2 py-1 bg-slate-100 rounded text-xs text-slate-600"
                            >
                              {Icon && <Icon className="w-3 h-3" />}
                              <span>{facility}</span>
                            </div>
                          );
                        })}
                      </div>

                      <div className="flex items-center justify-between">
                        <div>
                          <span className="text-3xl font-bold text-stone-900">
                            ${room.price}
                          </span>
                          <span className="text-slate-500 text-sm">/month</span>
                        </div>
                        <Button
                          onClick={() => onRoomClick(room.id)}
                          className="bg-stone-900 hover:bg-stone-800"
                        >
                          View Details
                        </Button>
                      </div>
                    </div>
                  </Card>
                </div>
              ))}
            </Slider>
          </div>
        </motion.div>
      </section>

      {/* Why Choose Us */}
      <section className="bg-slate-50 py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6, delay: 0.4 }}
          >
            <div className="text-center mb-12">
              <h2 className="text-4xl font-bold text-slate-900 mb-4">Why Choose LuxeStay?</h2>
              <p className="text-lg text-slate-600">Premium features for premium living</p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
              {[
                {
                  title: 'Verified Listings',
                  description: 'All properties are verified and inspected for quality',
                  icon: '✓'
                },
                {
                  title: 'Flexible Booking',
                  description: 'Monthly rentals with flexible terms and conditions',
                  icon: '📅'
                },
                {
                  title: '24/7 Support',
                  description: 'Round-the-clock customer support for all your needs',
                  icon: '💬'
                },
              ].map((feature, index) => (
                <Card key={index} className="bg-white p-8 text-center hover:shadow-lg transition-shadow">
                  <div className="text-5xl mb-4">{feature.icon}</div>
                  <h3 className="text-xl font-bold text-slate-900 mb-2">{feature.title}</h3>
                  <p className="text-slate-600">{feature.description}</p>
                </Card>
              ))}
            </div>
          </motion.div>
        </div>
      </section>
    </div>
  );
}