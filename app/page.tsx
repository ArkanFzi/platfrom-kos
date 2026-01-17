"use client";

import { useState } from 'react';
import { Button } from '@/app/components/ui/button';

// Login Component
import { Login } from '@/app/components/Login';

// Admin Components
import { AdminSidebar } from '@/app/components/admin/AdminSidebar';
import { LuxuryDashboard } from '@/app/components/admin/LuxuryDashboard';
import { GalleryData } from '@/app/components/admin/GalleryData';
import { LuxuryRoomManagement } from '@/app/components/admin/LuxuryRoomManagement';
import { TenantData } from '@/app/components/admin/TenantData';
import { LuxuryPaymentConfirmation } from '@/app/components/admin/LuxuryPaymentConfirmation';
import { LuxuryReports } from '@/app/components/admin/LuxuryReports';

// Tenant Components
import { UserPlatform } from '@/app/components/tenant/user-platform';

type ViewMode = 'login' | 'home' | 'admin' | 'tenant';
type AdminPage = 'dashboard' | 'gallery' | 'rooms' | 'tenants' | 'payments' | 'reports';
type TenantPage = 'landing' | 'room-detail' | 'booking' | 'payment' | 'history';

export default function App() {
  const [viewMode, setViewMode] = useState<ViewMode>('login');
  const [adminPage, setAdminPage] = useState<AdminPage>('dashboard');
  const [tenantPage, setTenantPage] = useState<TenantPage>('landing');
  const [selectedRoomId, setSelectedRoomId] = useState<string | null>(null);
  const [bookingData, setBookingData] = useState<any>(null);
  const [userRole, setUserRole] = useState<'admin' | 'tenant' | 'guest' | null>(null);

  // Login Screen
  if (viewMode === 'login') {
    return (
      <Login
        onLoginAsAdmin={() => {
          setUserRole('admin');
          setViewMode('admin');
        }}
        onLoginAsUser={() => {
          setUserRole('tenant');
          setViewMode('tenant');
        }}
        onLoginAsGuest={() => {
          setUserRole('guest');
          setViewMode('home');
        }}
      />
    );
  }

  // Home Selection Screen
  if (viewMode === 'home') {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-600 via-blue-700 to-blue-800 flex items-center justify-center p-4">
        <div className="max-w-4xl w-full">
          <div className="text-center mb-12">
            <h1 className="text-5xl font-bold text-white mb-4">
              Kos-kosan Management System
            </h1>
            <p className="text-xl text-blue-100">
              Complete boarding house management solution
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Admin Portal */}
            <div className="bg-white rounded-2xl shadow-2xl overflow-hidden hover:scale-105 transition-transform">
              <div className="bg-gradient-to-r from-slate-800 to-slate-900 p-8 text-white">
                <div className="size-16 bg-white/20 rounded-xl flex items-center justify-center mb-4">
                  <svg className="size-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                  </svg>
                </div>
                <h2 className="text-2xl font-semibold mb-2">Admin Portal</h2>
                <p className="text-blue-100">
                  Manage rooms, tenants, payments, and reports
                </p>
              </div>
              <div className="p-8">
                <ul className="space-y-3 mb-6 text-slate-600">
                  <li className="flex items-center gap-2">
                    <span className="size-1.5 bg-blue-600 rounded-full"></span>
                    Dashboard & Analytics
                  </li>
                  <li className="flex items-center gap-2">
                    <span className="size-1.5 bg-blue-600 rounded-full"></span>
                    Room Management (CRUD)
                  </li>
                  <li className="flex items-center gap-2">
                    <span className="size-1.5 bg-blue-600 rounded-full"></span>
                    Payment Confirmation
                  </li>
                  <li className="flex items-center gap-2">
                    <span className="size-1.5 bg-blue-600 rounded-full"></span>
                    Financial Reports
                  </li>
                </ul>
                <Button 
                  className="w-full" 
                  size="lg"
                  onClick={() => setViewMode('admin')}
                >
                  Enter Admin Portal
                </Button>
              </div>
            </div>

            {/* Tenant Portal */}
            <div className="bg-white rounded-2xl shadow-2xl overflow-hidden hover:scale-105 transition-transform">
              <div className="bg-gradient-to-r from-blue-600 to-blue-700 p-8 text-white">
                <div className="size-16 bg-white/20 rounded-xl flex items-center justify-center mb-4">
                  <svg className="size-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
                  </svg>
                </div>
                <h2 className="text-2xl font-semibold mb-2">Tenant Portal</h2>
                <p className="text-blue-100">
                  Browse rooms, book, and manage your rentals
                </p>
              </div>
              <div className="p-8">
                <ul className="space-y-3 mb-6 text-slate-600">
                  <li className="flex items-center gap-2">
                    <span className="size-1.5 bg-blue-600 rounded-full"></span>
                    Search & Browse Rooms
                  </li>
                  <li className="flex items-center gap-2">
                    <span className="size-1.5 bg-blue-600 rounded-full"></span>
                    View Room Details
                  </li>
                  <li className="flex items-center gap-2">
                    <span className="size-1.5 bg-blue-600 rounded-full"></span>
                    Online Booking & Payment
                  </li>
                  <li className="flex items-center gap-2">
                    <span className="size-1.5 bg-blue-600 rounded-full"></span>
                    Rental History
                  </li>
                </ul>
                <Button 
                  className="w-full" 
                  size="lg"
                  variant="outline"
                  onClick={() => setViewMode('tenant')}
                >
                  Enter Tenant Portal
                </Button>
              </div>
            </div>
          </div>

          <div className="text-center mt-12">
            <Button 
              variant="ghost" 
              className="text-blue-200 hover:text-white hover:bg-blue-600/20 mr-4"
              onClick={() => setViewMode('login')}
            >
              Back to Login
            </Button>
            <p className="text-blue-100 text-sm mt-4">
              © 2026 Kos-kosan Management System. All rights reserved.
            </p>
          </div>
        </div>
      </div>
    );
  }

  // Admin Portal
  if (viewMode === 'admin') {
    return (
      <div className="flex h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950">
        <AdminSidebar 
          currentPage={adminPage} 
          onNavigate={(page) => setAdminPage(page as AdminPage)} 
        />
        <div className="flex-1 overflow-auto">
          <div className="p-4 bg-slate-900/50 border-b border-slate-800/50 flex items-center justify-between">
            <Button 
              variant="outline" 
              onClick={() => setViewMode('home')}
              className="bg-slate-800/50 border-slate-700 text-white hover:bg-slate-800"
            >
              Back to Home
            </Button>
            <Button 
              variant="destructive" 
              onClick={() => setViewMode('login')}
              className="bg-red-600 hover:bg-red-700"
            >
              Logout
            </Button>
          </div>
          {adminPage === 'dashboard' && <LuxuryDashboard />}
          {adminPage === 'gallery' && <GalleryData />}
          {adminPage === 'rooms' && <LuxuryRoomManagement />}
          {adminPage === 'tenants' && <TenantData />}
          {adminPage === 'payments' && <LuxuryPaymentConfirmation />}
          {adminPage === 'reports' && <LuxuryReports />}
        </div>
      </div>
    );
  }

  // Tenant Portal
  if (viewMode === 'tenant') {
    return (
      <UserPlatform 
        onLogout={() => setViewMode('login')}
      />
    );
  }

  return null;
}

