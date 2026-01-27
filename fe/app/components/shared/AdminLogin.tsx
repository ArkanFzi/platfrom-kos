"use client";

import { useState } from 'react';
import { motion } from 'framer-motion';
import { Button } from '@/app/components/ui/button';
import { Input } from '@/app/components/ui/input';
import { Label } from '@/app/components/ui/label';
import { Shield, ArrowLeft } from 'lucide-react';
import { api } from '@/app/services/api';

interface AdminLoginProps {
  onLoginSuccess: () => void;
  onBack: () => void;
}

export function AdminLogin({ onLoginSuccess, onBack }: AdminLoginProps) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');
    try {
      const data = await api.login({ username, password });
      if (data.user.role !== 'admin') {
        api.logout();
        setError('Access denied. Admin privileges required.');
        return;
      }
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
    <div className="flex h-screen bg-slate-900 overflow-hidden">
      <div className="hidden md:flex md:w-1/2 bg-gradient-to-br from-slate-800 to-slate-950 items-center justify-center p-12 text-white border-r border-slate-800">
        <div className="text-center max-w-md">
          <div className="bg-amber-500/10 p-6 rounded-3xl inline-block mb-8">
            <Shield className="w-20 h-20 text-amber-500" />
          </div>
          <h1 className="text-5xl font-bold mb-6">Admin Panel</h1>
          <p className="text-xl text-slate-400">Secure access for system administrators only. Manage your property with ease.</p>
        </div>
      </div>

      <div className="flex w-full md:w-1/2 items-center justify-center p-8 bg-slate-950">
        <motion.div 
          initial={{ opacity: 0, x: 20 }}
          animate={{ opacity: 1, x: 0 }}
          className="w-full max-w-sm"
        >
          <Button variant="ghost" onClick={onBack} className="mb-8 p-0 text-slate-400 hover:text-white hover:bg-transparent">
            <ArrowLeft className="w-4 h-4 mr-2" /> Back to selection
          </Button>

          <h2 className="text-4xl font-bold text-white mb-2">Login Admin</h2>
          <p className="text-slate-400 mb-10">Verification required to proceed</p>

          {error && (
            <div className="bg-red-500/10 text-red-400 p-4 rounded-xl mb-8 text-sm font-medium border border-red-500/20">
              {error}
            </div>
          )}

          <form onSubmit={handleLogin} className="space-y-6">
            <div className="space-y-2">
              <Label htmlFor="username" className="text-slate-300">Admin Username</Label>
              <Input
                id="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Enter admin username"
                className="bg-slate-900 border-slate-700 text-white focus:border-amber-500"
                required
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="password" className="text-slate-300">Security Password</Label>
              <Input
                id="password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="••••••••"
                className="bg-slate-900 border-slate-700 text-white focus:border-amber-500"
                required
              />
            </div>

            <Button
              type="submit"
              disabled={isLoading}
              className="w-full bg-amber-600 hover:bg-amber-700 text-white h-12 font-bold text-lg rounded-xl mt-4 shadow-lg shadow-amber-600/20 transition-all hover:scale-[1.02]"
            >
              {isLoading ? 'Verifying...' : 'Authenticate'}
            </Button>
          </form>

          <div className="mt-12 pt-8 border-t border-slate-800 text-center">
            <p className="text-slate-500 text-sm">
              Protected by military-grade encryption.<br/>
              All login attempts are logged.
            </p>
          </div>
        </motion.div>
      </div>
    </div>
  );
}
