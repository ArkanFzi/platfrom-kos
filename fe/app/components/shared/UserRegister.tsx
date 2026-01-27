"use client";

import { useState } from 'react';
import { motion } from 'framer-motion';
import { Button } from '@/app/components/ui/button';
import { Input } from '@/app/components/ui/input';
import { Label } from '@/app/components/ui/label';
import { UserPlus, ArrowLeft, CheckCircle2 } from 'lucide-react';
import { api } from '@/app/services/api';
import { ImageWithFallback } from './ImageWithFallback';

interface UserRegisterProps {
  onRegisterSuccess: () => void;
  onBackToLogin: () => void;
}

export function UserRegister({ onRegisterSuccess, onBackToLogin }: UserRegisterProps) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');
  const [isSuccess, setIsSuccess] = useState(false);

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    if (password !== confirmPassword) {
      setError('Passwords do not match');
      return;
    }

    setIsLoading(true);
    try {
      await api.register({ username, password, role: 'tenant' });
      setIsSuccess(true);
      setTimeout(() => {
        onRegisterSuccess();
      }, 2000);
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError('Registration failed');
      }
    } finally {
      setIsLoading(false);
    }
  };

  if (isSuccess) {
    return (
      <div className="flex h-screen bg-slate-50 items-center justify-center p-8">
        <motion.div 
          initial={{ scale: 0.9, opacity: 0 }}
          animate={{ scale: 1, opacity: 1 }}
          className="bg-white p-12 rounded-3xl shadow-2xl text-center max-w-sm w-full border border-slate-100"
        >
          <div className="size-20 bg-emerald-100 rounded-full flex items-center justify-center mx-auto mb-6">
            <CheckCircle2 className="w-10 h-10 text-emerald-600" />
          </div>
          <h2 className="text-3xl font-bold text-slate-900 mb-2">Account Created!</h2>
          <p className="text-slate-600 mb-8">Registration successful. You are being redirected to login...</p>
        </motion.div>
      </div>
    );
  }

  return (
    <div className="flex h-screen bg-slate-50 overflow-hidden">
      <div className="hidden md:flex md:w-1/2 bg-stone-800 items-center justify-center p-12 text-white relative">
        <ImageWithFallback 
          src="https://images.unsplash.com/photo-1554995207-c18c20360a59?q=80&w=1000"
          alt="Registration Background"
          className="absolute inset-0 bg-cover bg-center opacity-30 w-full h-full object-cover"
        />
        <div className="relative z-10 text-center max-w-md">
          <UserPlus className="w-16 h-16 mx-auto mb-6 text-amber-400" />
          <h1 className="text-4xl font-bold mb-4">Start Your Journey</h1>
          <p className="text-xl text-stone-300">Join our premium community today. Find comfort, style, and convenience.</p>
        </div>
      </div>

      <div className="flex w-full md:w-1/2 items-center justify-center p-8">
        <motion.div 
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          className="w-full max-w-sm"
        >
          <Button variant="ghost" onClick={onBackToLogin} className="mb-8 p-0 hover:bg-transparent text-slate-500">
            <ArrowLeft className="w-4 h-4 mr-2" /> Back to Login
          </Button>

          <h2 className="text-3xl font-bold text-slate-900 mb-2">Register</h2>
          <p className="text-slate-600 mb-8">Create your tenant account</p>

          {error && (
            <div className="bg-red-50 text-red-600 p-3 rounded-lg mb-6 text-sm font-medium border border-red-100">
              {error}
            </div>
          )}

          <form onSubmit={handleRegister} className="space-y-4">
            <div>
              <Label htmlFor="username">Choose Username</Label>
              <Input
                id="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Ex: arkan_tenant"
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
            <div>
              <Label htmlFor="confirm-password">Confirm Password</Label>
              <Input
                id="confirm-password"
                type="password"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                placeholder="••••••••"
                className="mt-1"
                required
              />
            </div>

            <Button
              type="submit"
              disabled={isLoading}
              className="w-full bg-stone-900 hover:bg-stone-800 text-white h-11 font-semibold mt-6"
            >
              {isLoading ? 'Creating Account...' : 'Register Account'}
            </Button>
          </form>

          <p className="mt-8 text-center text-slate-600">
            Already have an account?{' '}
            <button 
              onClick={onBackToLogin}
              className="text-amber-600 font-semibold hover:underline"
            >
              Log in instead
            </button>
          </p>
        </motion.div>
      </div>
    </div>
  );
}
