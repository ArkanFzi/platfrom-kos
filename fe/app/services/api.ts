// Standard API Response Format yang dikirim backend
interface ApiResponse<T = any> {
  success: boolean;
  data?: T;
  message?: string;
  errors?: string[];
}

interface ApiError extends Error {
  status: number;
  errors?: string[];
}

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8081/api';

// Create custom error class
class ApiErrorClass extends Error implements ApiError {
  status: number;
  errors?: string[];

  constructor(message: string, status: number, errors?: string[]) {
    super(message);
    this.status = status;
    this.errors = errors;
    Object.setPrototypeOf(this, ApiErrorClass.prototype);
  }
}

// Parse standard API response
const parseApiResponse = async <T = any>(res: Response): Promise<T> => {
  const contentType = res.headers.get('content-type');
  
  if (!contentType || !contentType.includes('application/json')) {
    if (res.status === 401) {
      // Token invalid/expired, clear auth
      if (typeof window !== 'undefined') {
        localStorage.removeItem('user');
        localStorage.removeItem('app_view_mode');
      }
      throw new ApiErrorClass('Unauthorized: Please login again', 401);
    }
    throw new ApiErrorClass(`Server error ${res.status}: ${res.statusText}`, res.status);
  }

  const responseData: ApiResponse<T> = await res.json();

  // Check if response has success flag
  if (responseData.success === false) {
    const errorMessage = responseData.message || 'An error occurred';
    throw new ApiErrorClass(errorMessage, res.status, responseData.errors);
  }

  if (!res.ok && responseData.success !== true) {
    const errorMessage = responseData.message || `Request failed with status ${res.status}`;
    throw new ApiErrorClass(errorMessage, res.status, responseData.errors);
  }

  // Return data dari response
  return responseData.data || responseData;
};

// HTTP methods wrapper dengan proper error handling
const apiCall = async <T = any>(
  method: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH',
  endpoint: string,
  body?: FormData | Record<string, any> | null,
  customHeaders?: Record<string, string>
): Promise<T> => {
  const url = `${API_URL}${endpoint}`;
  const options: RequestInit = {
    method,
    credentials: 'include', // Include cookies in requests
    headers: {
      ...customHeaders,
    },
  };

  // Handle body
  if (body) {
    if (body instanceof FormData) {
      options.body = body;
      // Don't set Content-Type for FormData, browser will set it with boundary
    } else {
      options.body = JSON.stringify(body);
      if (!options.headers) options.headers = {};
      (options.headers as Record<string, string>)['Content-Type'] = 'application/json';
    }
  }

  try {
    const res = await fetch(url, options);
    return await parseApiResponse<T>(res);
  } catch (error) {
    if (error instanceof ApiErrorClass) {
      throw error;
    }
    // Network error atau parsing error
    const message = error instanceof Error ? error.message : 'Unknown error occurred';
    throw new ApiErrorClass(message, 0);
  }
};

// API endpoints
export const api = {
  // ============ AUTH ============
  login: async (credentials: { username: string; password: string }) => {
    const response = await apiCall<{
      access_token: string;
      user: { id: number; username: string; email: string; role: string };
    }>('POST', '/auth/login', credentials);
    
    // Token sekarang di httpOnly cookie, hanya save user data
    if (typeof window !== 'undefined' && response.user) {
      localStorage.setItem('user', JSON.stringify(response.user));
    }
    return response;
  },

  register: async (data: {
    username: string;
    email: string;
    password: string;
    nama_lengkap: string;
    no_telepon: string;
  }) => {
    return apiCall('POST', '/auth/register', data);
  },

  logout: async () => {
    // Clear local user data
    if (typeof window !== 'undefined') {
      localStorage.removeItem('user');
      localStorage.removeItem('app_view_mode');
    }
    // Backend akan clear cookie automatically
    return { success: true };
  },

  // ============ PROFILE ============
  getProfile: async () => {
    return apiCall('GET', '/profile');
  },

  updateProfile: async (data: Record<string, any>) => {
    return apiCall('PUT', '/profile', data);
  },

  changePassword: async (data: { old_password: string; new_password: string }) => {
    return apiCall('PUT', '/profile/change-password', data);
  },

  // ============ ROOMS ============
  getRooms: async () => {
    return apiCall('GET', '/kamar');
  },

  getRoomById: async (id: string) => {
    return apiCall('GET', `/kamar/${id}`);
  },

  createRoom: async (formData: FormData) => {
    return apiCall('POST', '/kamar', formData);
  },

  updateRoom: async (id: string, formData: FormData) => {
    return apiCall('PUT', `/kamar/${id}`, formData);
  },

  deleteRoom: async (id: string) => {
    return apiCall('DELETE', `/kamar/${id}`);
  },

  // ============ GALLERY ============
  getGalleries: async () => {
    return apiCall('GET', '/galleries');
  },

  createGallery: async (formData: FormData) => {
    return apiCall('POST', '/galleries', formData);
  },

  deleteGallery: async (id: string) => {
    return apiCall('DELETE', `/galleries/${id}`);
  },

  // ============ BOOKINGS ============
  getMyBookings: async () => {
    return apiCall('GET', '/bookings');
  },

  createBooking: async (data: {
    kamar_id: number;
    tanggal_mulai: string;
    durasi_sewa: number;
  }) => {
    return apiCall('POST', '/bookings', data);
  },

  // ============ PAYMENTS ============
  createSnapToken: async (pemesananId: number) => {
    return apiCall('POST', '/payments/snap-token', { pemesanan_id: pemesananId });
  },

  confirmCashPayment: async (bookingId: number) => {
    return apiCall('POST', `/payments/confirm-cash/${bookingId}`);
  },

  // ============ REVIEWS ============
  getReviews: async (roomId: string) => {
    return apiCall('GET', `/kamar/${roomId}/reviews`);
  },

  createReview: async (data: {
    kamar_id: number;
    rating: number;
    comment: string;
  }) => {
    return apiCall('POST', '/reviews', data);
  },

  // ============ DASHBOARD ============
  getDashboardStats: async () => {
    return apiCall('GET', '/dashboard');
  },

  // ============ ADMIN - PAYMENTS ============
  getAllPayments: async () => {
    return apiCall('GET', '/payments');
  },

  confirmPayment: async (paymentId: string) => {
    return apiCall('PUT', `/payments/${paymentId}/confirm`);
  },

  // ============ ADMIN - TENANTS ============
  getAllTenants: async () => {
    return apiCall('GET', '/tenants');
  },

  // ============ CONTACT ============
  submitContactForm: async (data: {
    name: string;
    email: string;
    subject: string;
    message: string;
  }) => {
    return apiCall('POST', '/contact', data);
  },

  // ============ HEALTH CHECK ============
  healthCheck: async () => {
    return apiCall('GET', '/health');
  },
};

export type { ApiResponse, ApiError };
export { ApiErrorClass };
