import { useEffect, useState } from 'react';
import { Users, UserCheck, Search, Loader2 } from 'lucide-react';
import { api } from '@/app/services/api';
import { Input } from '@/app/components/ui/input';
import type { Tenant } from '@/app/services/api';

export function UserManagement() {
  const [allUsers, setAllUsers] = useState<Tenant[]>([]);
  const [filteredUsers, setFilteredUsers] = useState<Tenant[]>([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [isLoading, setIsLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<'all' | 'guest' | 'tenant'>('all');

  useEffect(() => {
    const fetchUsers = async () => {
      setIsLoading(true);
      try {
        const data = await api.getAllTenants();
        setAllUsers(data);
        setFilteredUsers(data);
      } catch (e) {
        console.error(e);
      } finally {
        setIsLoading(false);
      }
    };
    void fetchUsers();
  }, []);

  useEffect(() => {
    let filtered = allUsers;
    
    // Filter by role based on active tab
    if (activeTab === 'guest') {
      filtered = filtered.filter(u => u.role === 'guest');
    } else if (activeTab === 'tenant') {
      filtered = filtered.filter(u => u.role === 'tenant');
    }

    // Filter by search query
    if (searchQuery) {
      filtered = filtered.filter(user =>
        user.nama_lengkap?.toLowerCase().includes(searchQuery.toLowerCase()) ||
        user.email?.toLowerCase().includes(searchQuery.toLowerCase()) ||
        user.nomor_hp?.includes(searchQuery)
      );
    }

    setFilteredUsers(filtered);
  }, [activeTab, searchQuery, allUsers]);

  const guestCount = allUsers.filter(u => u.role === 'guest').length;
  const tenantCount = allUsers.filter(u => u.role === 'tenant').length;

  const getRoleBadge = (role?: string) => {
    if (role === 'tenant') {
      return (
        <span className="px-3 py-1 bg-green-500/10 text-green-400 text-[10px] font-bold rounded-lg border border-green-500/20 uppercase">
          Tenant
        </span>
      );
    }
    return (
      <span className="px-3 py-1 bg-amber-500/10 text-amber-400 text-[10px] font-bold rounded-lg border border-amber-500/20 uppercase">
        Guest
      </span>
    );
  };

  return (
    <div className="p-4 md:p-8 space-y-6 md:space-y-8 bg-slate-950 min-h-screen">
      {/* Header */}
      <div className="flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h2 className="text-2xl md:text-3xl font-bold text-amber-500">User Management</h2>
          <p className="text-slate-400 text-xs md:text-sm">Kelola data guest dan tenant</p>
        </div>

        {/* Search Bar */}
        <div className="relative w-full md:w-96">
          <Search className="absolute left-4 top-1/2 -translate-y-1/2 size-4 text-slate-500" />
          <Input
            placeholder="Cari user..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="pl-12 bg-slate-900 border-slate-800 text-white placeholder:text-slate-500 h-11 md:h-12 rounded-xl"
          />
        </div>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="bg-gradient-to-br from-slate-800/50 to-slate-900/50 backdrop-blur-sm rounded-2xl p-6 border border-slate-700/50">
          <div className="flex items-center gap-4">
            <div className="size-12 bg-blue-500/10 rounded-xl flex items-center justify-center">
              <Users className="size-6 text-blue-400" />
            </div>
            <div>
              <p className="text-slate-400 text-xs uppercase font-bold tracking-widest">Total Users</p>
              <p className="text-2xl font-bold text-white">{allUsers.length}</p>
            </div>
          </div>
        </div>

        <div className="bg-gradient-to-br from-amber-800/20 to-amber-900/20 backdrop-blur-sm rounded-2xl p-6 border border-amber-700/30">
          <div className="flex items-center gap-4">
            <div className="size-12 bg-amber-500/10 rounded-xl flex items-center justify-center">
              <Users className="size-6 text-amber-400" />
            </div>
            <div>
              <p className="text-slate-400 text-xs uppercase font-bold tracking-widest">Guest Users</p>
              <p className="text-2xl font-bold text-white">{guestCount}</p>
            </div>
          </div>
        </div>

        <div className="bg-gradient-to-br from-green-800/20 to-green-900/20 backdrop-blur-sm rounded-2xl p-6 border border-green-700/30">
          <div className="flex items-center gap-4">
            <div className="size-12 bg-green-500/10 rounded-xl flex items-center justify-center">
              <UserCheck className="size-6 text-green-400" />
            </div>
            <div>
              <p className="text-slate-400 text-xs uppercase font-bold tracking-widest">Active Tenants</p>
              <p className="text-2xl font-bold text-white">{tenantCount}</p>
            </div>
          </div>
        </div>
      </div>

      {/* Tabs */}
      <div className="flex gap-2 border-b border-slate-800 overflow-x-auto">
        <button
          onClick={() => setActiveTab('all')}
          className={`px-6 py-3 font-bold text-sm transition-all whitespace-nowrap ${
            activeTab === 'all'
              ? 'text-amber-500 border-b-2 border-amber-500'
              : 'text-slate-500 hover:text-slate-300'
          }`}
        >
          All Users ({allUsers.length})
        </button>
        <button
          onClick={() => setActiveTab('guest')}
          className={`px-6 py-3 font-bold text-sm transition-all whitespace-nowrap ${
            activeTab === 'guest'
              ? 'text-amber-500 border-b-2 border-amber-500'
              : 'text-slate-500 hover:text-slate-300'
          }`}
        >
          Guest Users ({guestCount})
        </button>
        <button
          onClick={() => setActiveTab('tenant')}
          className={`px-6 py-3 font-bold text-sm transition-all whitespace-nowrap ${
            activeTab === 'tenant'
              ? 'text-amber-500 border-b-2 border-amber-500'
              : 'text-slate-500 hover:text-slate-300'
          }`}
        >
          Active Tenants ({tenantCount})
        </button>
      </div>

      {/* Content Container */}
      <div className="bg-slate-900/20 md:border md:border-slate-800 rounded-[1.5rem] md:rounded-[2rem] p-4 md:p-8 backdrop-blur-sm">
        {/* Desktop Header */}
        <div className="hidden md:grid grid-cols-[80px_1.5fr_1fr_1.5fr_120px] gap-4 px-6 mb-8 text-slate-500 text-xs font-bold uppercase tracking-widest">
          <div>ID</div>
          <div>Nama Lengkap</div>
          <div>NIK</div>
          <div>Kontak</div>
          <div className="text-center">Status</div>
        </div>

        {/* Data Rows */}
        <div className="space-y-4">
          {isLoading ? (
            <div className="flex flex-col items-center justify-center py-20 gap-4">
              <Loader2 className="size-10 text-amber-500 animate-spin" />
              <p className="text-slate-500 italic">Mengambil data user...</p>
            </div>
          ) : filteredUsers.length === 0 ? (
            <div className="text-center py-20 bg-slate-800/20 rounded-2xl border border-dashed border-slate-800">
              <Search className="size-12 text-slate-700 mx-auto mb-4" />
              <p className="text-slate-400 font-medium">Tidak ada user yang cocok dengan kriteria</p>
            </div>
          ) : (
            filteredUsers.map((user) => (
              <div key={user.id}>
                {/* Desktop Row */}
                <div className="hidden md:grid grid-cols-[80px_1.5fr_1fr_1.5fr_120px] gap-4 px-6 py-4 items-center bg-slate-800/20 hover:bg-slate-800/40 rounded-2xl transition-all border border-transparent hover:border-slate-800 group">
                  <div className="text-slate-500 font-mono text-sm">#{user.id}</div>
                  <div className="text-white font-bold">{user.nama_lengkap || 'N/A'}</div>
                  <div className="text-slate-300 text-sm font-mono">{user.nik || '-'}</div>
                  <div className="flex flex-col">
                    <span className="text-slate-300 text-sm truncate">{user.email || 'N/A'}</span>
                    <span className="text-slate-500 text-xs">{user.nomor_hp || '-'}</span>
                  </div>
                  <div className="flex justify-center">
                    {getRoleBadge(user.role)}
                  </div>
                </div>

                {/* Mobile Card */}
                <div className="md:hidden p-5 bg-slate-900/40 border border-slate-800 rounded-2xl space-y-4">
                  <div className="flex justify-between items-start">
                    <div className="flex items-center gap-3">
                      <div className="size-10 bg-gradient-to-br from-amber-500 to-amber-600 rounded-xl flex items-center justify-center font-bold text-white">
                        {user.nama_lengkap?.charAt(0) || 'U'}
                      </div>
                      <div>
                        <h3 className="text-white font-bold">{user.nama_lengkap || 'N/A'}</h3>
                        <p className="text-slate-500 font-mono text-[10px]">ID: #{user.id}</p>
                        <p className="text-slate-500 font-mono text-[10px]">NIK: {user.nik || '-'}</p>
                      </div>
                    </div>
                    {getRoleBadge(user.role)}
                  </div>

                  <div className="pt-2 border-t border-slate-800">
                    <p className="text-slate-500 text-[10px] uppercase font-bold tracking-widest mb-1">Email / Telepon</p>
                    <p className="text-slate-300 text-xs truncate mb-1">{user.email || 'N/A'}</p>
                    <p className="text-slate-500 text-xs">{user.nomor_hp || '-'}</p>
                  </div>
                </div>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
}
