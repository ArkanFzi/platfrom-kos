import { useState } from 'react';
import { Card } from '@/app/components/ui/card';
import { Button } from '@/app/components/ui/button';
import { Input } from '@/app/components/ui/input';
import { Label } from '@/app/components/ui/label';
import { Progress } from '@/app/components/ui/progress';
import { 
  ArrowLeft, 
  Check, 
  CreditCard, 
  Upload,
  CheckCircle2
} from 'lucide-react';

interface BookingFlowProps {
  roomId: string;
  onBack: () => void;
}

export function BookingFlow({ roomId, onBack }: BookingFlowProps) {
  const [step, setStep] = useState(1);
  const [bookingComplete, setBookingComplete] = useState(false);
  const [formData, setFormData] = useState({
    fullName: '',
    email: '',
    phone: '',
    moveInDate: '',
    duration: '6',
    paymentMethod: 'bank',
  });

  const steps = [
    { number: 1, title: 'Personal Info', description: 'Your details' },
    { number: 2, title: 'Booking Details', description: 'Move-in & duration' },
    { number: 3, title: 'Payment', description: 'Complete payment' },
  ];

  const handleInputChange = (field: string, value: string) => {
    setFormData({ ...formData, [field]: value });
  };

  const nextStep = () => {
    if (step < 3) {
      setStep(step + 1);
    } else {
      setBookingComplete(true);
    }
  };

  const prevStep = () => {
    if (step > 1) {
      setStep(step - 1);
    }
  };

  const progress = (step / 3) * 100;

  if (bookingComplete) {
    return (
      <div className="min-h-screen bg-slate-50 py-12 flex items-center justify-center">
        <div className="max-w-2xl w-full mx-4">
          <Card className="p-12 text-center">
            <div className="w-20 h-20 bg-emerald-100 rounded-full flex items-center justify-center mx-auto mb-6">
              <CheckCircle2 className="w-12 h-12 text-emerald-600" />
            </div>
            <h1 className="text-3xl font-bold text-slate-900 mb-4">Booking Confirmed!</h1>
            <p className="text-lg text-slate-600 mb-8">
              Your booking request has been submitted successfully. We'll review your payment and send a confirmation email within 24 hours.
            </p>
            <div className="space-y-3 p-6 bg-slate-50 rounded-lg mb-8">
              <div className="flex justify-between">
                <span className="text-slate-600">Booking ID:</span>
                <span className="font-medium">#BK{Date.now().toString().slice(-6)}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-slate-600">Name:</span>
                <span className="font-medium">{formData.fullName}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-slate-600">Email:</span>
                <span className="font-medium">{formData.email}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-slate-600">Move-in Date:</span>
                <span className="font-medium">{formData.moveInDate}</span>
              </div>
            </div>
            <Button onClick={onBack} className="w-full bg-stone-900 hover:bg-stone-800">
              Back to Homepage
            </Button>
          </Card>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-slate-50 py-8">
      <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Back Button */}
        <Button variant="outline" onClick={onBack} className="mb-6">
          <ArrowLeft className="w-4 h-4 mr-2" />
          Back
        </Button>

        {/* Progress Steps */}
        <Card className="p-6 mb-8">
          <div className="flex items-center justify-between mb-4">
            {steps.map((s, index) => (
              <div key={s.number} className="flex items-center flex-1">
                <div className="flex flex-col items-center">
                  <div
                    className={`w-12 h-12 rounded-full flex items-center justify-center font-bold transition-all ${
                      step >= s.number
                        ? 'bg-stone-900 text-white'
                        : 'bg-slate-200 text-slate-400'
                    }`}
                  >
                    {step > s.number ? <Check className="w-6 h-6" /> : s.number}
                  </div>
                  <div className="mt-2 text-center">
                    <p className="text-sm font-medium text-slate-900">{s.title}</p>
                    <p className="text-xs text-slate-500">{s.description}</p>
                  </div>
                </div>
                {index < steps.length - 1 && (
                  <div
                    className={`flex-1 h-1 mx-4 rounded transition-all ${
                      step > s.number ? 'bg-stone-900' : 'bg-slate-200'
                    }`}
                  />
                )}
              </div>
            ))}
          </div>
          <Progress value={progress} className="h-2" />
        </Card>

        {/* Step Content */}
        <div>
          <Card className="p-8">
            {step === 1 && (
              <div className="space-y-6">
                <div>
                  <h2 className="text-2xl font-bold text-slate-900 mb-2">Personal Information</h2>
                  <p className="text-slate-600">Tell us about yourself</p>
                </div>

                <div className="space-y-4">
                  <div>
                    <Label htmlFor="fullName">Full Name *</Label>
                    <Input
                      id="fullName"
                      placeholder="John Doe"
                      value={formData.fullName}
                      onChange={(e) => handleInputChange('fullName', e.target.value)}
                      className="mt-1"
                    />
                  </div>

                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div>
                      <Label htmlFor="email">Email Address *</Label>
                      <Input
                        id="email"
                        type="email"
                        placeholder="john@example.com"
                        value={formData.email}
                        onChange={(e) => handleInputChange('email', e.target.value)}
                        className="mt-1"
                      />
                    </div>
                    <div>
                      <Label htmlFor="phone">Phone Number *</Label>
                      <Input
                        id="phone"
                        type="tel"
                        placeholder="+1 (555) 123-4567"
                        value={formData.phone}
                        onChange={(e) => handleInputChange('phone', e.target.value)}
                        className="mt-1"
                      />
                    </div>
                  </div>
                </div>
              </div>
            )}

            {step === 2 && (
              <div className="space-y-6">
                <div>
                  <h2 className="text-2xl font-bold text-slate-900 mb-2">Booking Details</h2>
                  <p className="text-slate-600">Choose your move-in date and duration</p>
                </div>

                <div className="space-y-4">
                  <div>
                    <Label htmlFor="moveInDate">Preferred Move-in Date *</Label>
                    <Input
                      id="moveInDate"
                      type="date"
                      value={formData.moveInDate}
                      onChange={(e) => handleInputChange('moveInDate', e.target.value)}
                      className="mt-1"
                    />
                  </div>

                  <div>
                    <Label htmlFor="duration">Rental Duration *</Label>
                    <select
                      id="duration"
                      value={formData.duration}
                      onChange={(e) => handleInputChange('duration', e.target.value)}
                      className="w-full mt-1 px-4 py-2 border border-slate-300 rounded-lg"
                    >
                      <option value="3">3 months</option>
                      <option value="6">6 months</option>
                      <option value="12">12 months</option>
                    </select>
                  </div>

                  <div className="p-4 bg-blue-50 rounded-lg border border-blue-200">
                    <h4 className="font-medium text-blue-900 mb-2">Booking Summary</h4>
                    <div className="space-y-1 text-sm text-blue-800">
                      <p>Duration: {formData.duration} months</p>
                      <p>Monthly rent: $1,200</p>
                      <p>Security deposit: $1,200</p>
                      <p className="font-bold pt-2 border-t border-blue-200">Total due: $2,400</p>
                    </div>
                  </div>
                </div>
              </div>
            )}

            {step === 3 && (
              <div className="space-y-6">
                <div>
                  <h2 className="text-2xl font-bold text-slate-900 mb-2">Payment</h2>
                  <p className="text-slate-600">Complete your booking payment</p>
                </div>

                <div className="space-y-4">
                  <div>
                    <Label>Payment Method *</Label>
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-2">
                      <button
                        onClick={() => handleInputChange('paymentMethod', 'bank')}
                        className={`p-4 border-2 rounded-lg text-left transition-all ${
                          formData.paymentMethod === 'bank'
                            ? 'border-stone-900 bg-stone-50'
                            : 'border-slate-200 hover:border-slate-300'
                        }`}
                      >
                        <CreditCard className="w-6 h-6 mb-2" />
                        <p className="font-medium">Bank Transfer</p>
                        <p className="text-sm text-slate-500">Direct bank transfer</p>
                      </button>
                      <button
                        onClick={() => handleInputChange('paymentMethod', 'ewallet')}
                        className={`p-4 border-2 rounded-lg text-left transition-all ${
                          formData.paymentMethod === 'ewallet'
                            ? 'border-stone-900 bg-stone-50'
                            : 'border-slate-200 hover:border-slate-300'
                        }`}
                      >
                        <CreditCard className="w-6 h-6 mb-2" />
                        <p className="font-medium">E-Wallet</p>
                        <p className="text-sm text-slate-500">PayPal, Venmo, etc.</p>
                      </button>
                    </div>
                  </div>

                  {formData.paymentMethod === 'bank' && (
                    <div className="p-4 bg-slate-50 rounded-lg">
                      <h4 className="font-medium mb-2">Bank Account Details</h4>
                      <div className="space-y-1 text-sm text-slate-600">
                        <p>Bank: LuxeStay Bank</p>
                        <p>Account Number: 1234567890</p>
                        <p>Account Name: LuxeStay Properties</p>
                        <p>Amount: $2,400</p>
                      </div>
                    </div>
                  )}

                  <div>
                    <Label htmlFor="receipt">Upload Payment Receipt *</Label>
                    <div className="mt-2 border-2 border-dashed border-slate-300 rounded-lg p-8 text-center hover:border-stone-400 transition-colors cursor-pointer">
                      <Upload className="w-8 h-8 mx-auto mb-2 text-slate-400" />
                      <p className="text-sm text-slate-600">Click to upload or drag and drop</p>
                      <p className="text-xs text-slate-500 mt-1">PNG, JPG up to 5MB</p>
                    </div>
                  </div>

                  <div className="p-4 bg-amber-50 rounded-lg border border-amber-200">
                    <p className="text-sm text-amber-800">
                      ⚠️ Please upload a clear screenshot or photo of your payment receipt. Our team will verify the payment within 24 hours.
                    </p>
                  </div>
                </div>
              </div>
            )}

            {/* Navigation Buttons */}
            <div className="flex gap-4 mt-8 pt-6 border-t border-slate-200">
              <Button
                variant="outline"
                onClick={prevStep}
                disabled={step === 1}
                className="flex-1"
              >
                Previous
              </Button>
              <Button
                onClick={nextStep}
                className="flex-1 bg-stone-900 hover:bg-stone-800"
              >
                {step === 3 ? 'Complete Booking' : 'Next Step'}
              </Button>
            </div>
          </Card>
        </div>
      </div>
    </div>
  );
}
