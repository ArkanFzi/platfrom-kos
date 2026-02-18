# Frontend Architecture

Detail arsitektur frontend yang dibangun dengan **Next.js 16 (App Router)**, **React 18**, **SWR**, dan **Tailwind CSS**.

## Folder Structure

```
fe/app/
├── page.tsx                       # Homepage (public)
├── layout.tsx                     # Root layout (providers, global styles)
├── globals.css                    # Global CSS + Tailwind
├── login/page.tsx                 # Halaman login
├── register/page.tsx              # Halaman registrasi
├── reset-password/page.tsx        # Reset password
├── admin/page.tsx                 # Admin dashboard
│
├── components/
│   ├── admin/                     # Komponen khusus admin
│   │   ├── LuxuryDashboard.tsx       # Dashboard analytics
│   │   ├── LuxuryRoomManagement.tsx  # CRUD kamar
│   │   ├── LuxuryPaymentConfirmation.tsx  # Konfirmasi pembayaran
│   │   ├── LuxuryReports.tsx         # Laporan keuangan
│   │   ├── AdminSidebar.tsx          # Sidebar navigasi admin
│   │   ├── TenantData.tsx            # Data penyewa
│   │   ├── GalleryData.tsx           # Manajemen galeri
│   │   └── UserManagement.tsx        # Manajemen user
│   │
│   ├── tenant/                    # Komponen penyewa/public
│   │   ├── home/                     # Homepage sections
│   │   ├── booking/                  # Booking flow
│   │   ├── dashboard/                # Tenant dashboard & history
│   │   └── room-detail/             # Detail kamar
│   │
│   ├── shared/                    # Komponen reusable
│   ├── ui/                        # Shadcn UI components (25+ komponen)
│   └── providers/                 # Context providers
│
├── context/                       # State management
│   ├── AppContext.tsx                # Auth state, user data
│   ├── ThemeContext.tsx              # Dark/light theme
│   ├── types.ts                     # Type definitions
│   ├── useApp.ts                    # Custom hook untuk AppContext
│   └── useTheme.ts                  # Custom hook untuk theme
│
├── services/                      # API layer
│   ├── api.ts                        # Semua API calls (473 lines)
│   └── authgoogle.ts                 # Google OAuth helper
│
├── styles/                        # CSS modules
└── utils/                         # Helper functions
```

## API Client Layer (`api.ts`)

File `api.ts` adalah **single source of truth** untuk semua komunikasi dengan backend. Berisi:

1. **Interface definitions** — TypeScript types untuk semua entity
2. **`apiCall` helper** — Wrapper untuk `fetch` dengan HttpOnly cookie dan auto-refresh
3. **API object** — Fungsi untuk setiap endpoint

```typescript
// Dari fe/app/services/api.ts — Core helper

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8081/api';

const apiCall = async <T>(method: string, endpoint: string, body?: unknown): Promise<T> => {
    const config: RequestInit = {
        method,
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include', // Kirim HttpOnly cookies otomatis
    };

    // FormData support (untuk upload file)
    if (body instanceof FormData) {
        delete (config.headers as Record<string, string>)['Content-Type'];
        config.body = body;
    } else if (body) {
        config.body = JSON.stringify(body);
    }

    const res = await fetch(`${API_URL}${endpoint}`, config);

    // Auto-refresh token pada 401
    if (res.status === 401) {
        const refreshed = await refreshAccessToken();
        if (refreshed) return apiCall<T>(method, endpoint, body);
        // Redirect ke login jika refresh gagal
    }

    return safeJson(res);
};
```

Sumber: [`fe/app/services/api.ts`](file:///c:/Users/Arkan/Documents/coding/platfrom-kos/fe/app/services/api.ts)

### Penggunaan API

```typescript
// Di component manapun:
import { api } from '@/app/services/api';

// Ambil data kamar
const rooms = await api.getRooms();

// Login
const result = await api.login({ username: 'admin', password: 'pass' });

// Upload bukti bayar
await api.uploadPaymentProof(paymentId, file);

// Logout
await api.logout();
```

## State Management: Context API

### AppContext

Menyimpan state autentikasi dan data user secara global:

```typescript
// Dari fe/app/context/types.ts — Tipe yang digunakan

interface AppContextType {
    user: User | null;
    isLoading: boolean;
    login: (credentials: LoginCredentials) => Promise<void>;
    logout: () => Promise<void>;
    updateProfile: (data: Partial<Tenant>) => Promise<void>;
}
```

### Penggunaan Context

```typescript
// Di component:
import { useApp } from '@/app/context/useApp';

function MyComponent() {
    const { user, logout, isLoading } = useApp();

    if (isLoading) return <LoadingSkeleton />;
    if (!user) return <LoginPrompt />;

    return <div>Welcome, {user.username}!</div>;
}
```

### ThemeContext

Mengelola dark/light mode:

```typescript
import { useTheme } from '@/app/context/useTheme';

function ThemeToggle() {
    const { theme, toggleTheme } = useTheme();
    return <button onClick={toggleTheme}>{theme === 'dark' ? '☀️' : '🌙'}</button>;
}
```

## Data Fetching: SWR Pattern

SWR (Stale-While-Revalidate) digunakan untuk data yang sering berubah, memberikan pengalaman "instant" karena menampilkan data dari cache dulu lalu memperbarui di background:

```typescript
import useSWR from 'swr';
import { api } from '@/app/services/api';

function RoomList() {
    const { data: rooms, error, isLoading } = useSWR('rooms', api.getRooms);

    if (isLoading) return <Skeleton />;
    if (error) return <ErrorMessage />;

    return (
        <div className="grid grid-cols-3 gap-4">
            {rooms?.map(room => <RoomCard key={room.id} room={room} />)}
        </div>
    );
}
```

## UI Components: Shadcn UI

Menggunakan **25+ komponen** dari Shadcn UI (berbasis Radix UI):

| Kategori | Komponen |
|----------|----------|
| **Form** | Dialog, Select, Checkbox, Radio, Label, Input |
| **Navigation** | Tabs, NavigationMenu, Menubar, DropdownMenu |
| **Feedback** | AlertDialog, Tooltip, HoverCard, Progress |
| **Layout** | ScrollArea, Separator, Accordion, Popover |
| **Data** | Table (custom), Calendar, Avatar |
| **Overlay** | Dialog, Sheet (Vaul), ContextMenu |

## Route Protection: Next.js Middleware

```typescript
// Dari fe/middleware.ts
// Middleware yang redirect user tidak terautentikasi dari halaman protected
```

File `middleware.ts` di root frontend memeriksa keberadaan cookie autentikasi dan melakukan redirect sebelum halaman di-render.

## Komponen Admin vs Tenant

| Aspek | Admin Components | Tenant Components |
|-------|-----------------|-------------------|
| **Lokasi** | `components/admin/` | `components/tenant/` |
| **Akses** | Role `admin` saja | Semua user terautentikasi |
| **Contoh** | `LuxuryDashboard`, `LuxuryRoomManagement` | `BookingFlow`, `RoomDetail` |
| **Bundle** | Hanya di-load di `/admin` | Di-load di halaman publik |

---

> [!NOTE]
> Next.js secara otomatis melakukan **code splitting** berdasarkan route. Komponen admin tidak akan ter-bundle ke halaman publik, menjaga performa untuk user biasa.
