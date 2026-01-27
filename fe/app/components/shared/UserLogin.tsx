"use client";

import { useState } from 'react';
import { motion } from 'framer-motion';
import { Button } from '@/app/components/ui/button';
import { Input } from '@/app/components/ui/input';
import { Label } from '@/app/components/ui/label';
import { Home, ArrowLeft } from 'lucide-react';
import { api } from '@/app/services/api';
import { ImageWithFallback } from './ImageWithFallback';

interface UserLoginProps {
// ... existing props ...
  onLoginSuccess: () => void;
  onBack: () => void;
  onRegisterClick: () => void;
}

export function UserLogin({ onLoginSuccess, onBack, onRegisterClick }: UserLoginProps) {
// ... existing state ...
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');
    try {
      await api.login({ username, password });
      onLoginSuccess();
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError('Login failed');
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="flex h-screen bg-slate-50 overflow-hidden">
      <div className="hidden md:flex md:w-1/2 bg-stone-900 items-center justify-center p-12 text-white relative">
        <ImageWithFallback 
          src="https://images.unsplash.com/photo-1522708323590-d24dbb6b0267?q=80&w=1000"
          alt="Login Background"
          className="absolute inset-0 bg-cover bg-center opacity-40 w-full h-full object-cover"
          priority
        />
        <div className="relative z-10 text-center max-w-md">
          <Home className="w-16 h-16 mx-auto mb-6 text-amber-400" />
          <h1 className="text-4xl font-bold mb-4">Welcome Tenant</h1>
          <p className="text-xl text-stone-300">Log in to manage your bookings and find your perfect home.</p>
        </div>
      </div>

      <div className="flex w-full md:w-1/2 items-center justify-center p-8">
        <motion.div 
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          className="w-full max-w-sm"
        >
          <Button variant="ghost" onClick={onBack} className="mb-8 p-0 hover:bg-transparent">
            <ArrowLeft className="w-4 h-4 mr-2" /> Back
          </Button>

          <h2 className="text-3xl font-bold text-slate-900 mb-2">Tenant Login</h2>
          <p className="text-slate-600 mb-8">Enter your credentials to access your portal</p>

          {error && (
            <div className="bg-red-50 text-red-600 p-3 rounded-lg mb-6 text-sm font-medium border border-red-100">
              {error}
            </div>
          )}

          <form onSubmit={handleLogin} className="space-y-5">
            <div>
              <Label htmlFor="username">Username</Label>
              <Input
                id="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Enter your username"
                className="mt-1"
                required
              />
            </div>
            <div>
              <Label htmlFor="password">Password</Label>
              <Input
                id="password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="••••••••"
                className="mt-1"
                required
              />
            </div>

            <Button
              type="submit"
              disabled={isLoading}
              className="w-full bg-stone-900 hover:bg-stone-800 text-white h-11 font-semibold mt-4"
            >
              {isLoading ? 'Signing in...' : 'Sign In'}
            </Button>
          </form>

          <p className="mt-8 text-center text-slate-600">
            Don&apos;t have an account?{' '}
            <button 
              onClick={onRegisterClick}
              className="text-amber-600 font-semibold hover:underline"
            >
              Register here
            </button>
          </p>
        </motion.div>
      </div>
    </div>
  );
}
