# Google OAuth Setup Guide

## How to Enable Google Login in LuxeStay

### Step 1: Create Google Cloud Project
1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project (Name it "LuxeStay" or similar)
3. Enable the "Google+ API"

### Step 2: Create OAuth 2.0 Credentials
1. Go to "Credentials" in the left menu
2. Click "Create Credentials" → "OAuth client ID"
3. Select "Web application"
4. Add authorized redirect URIs:
   - `http://localhost:3000/api/auth/callback/google` (development)
   - `http://localhost:3000` (development)
   - `https://yourdomain.com/api/auth/callback/google` (production)
   - `https://yourdomain.com` (production)

5. Copy your **Client ID** and **Client Secret**

### Step 3: Install NextAuth.js (Optional but Recommended)
```bash
npm install next-auth @react-oauth/google
```

### Step 4: Add Environment Variables
Create a `.env.local` file in your project root:

```env
GOOGLE_CLIENT_ID=your_client_id_here
GOOGLE_CLIENT_SECRET=your_client_secret_here
NEXTAUTH_URL=http://localhost:3000
NEXTAUTH_SECRET=your_random_secret_here
```

To generate `NEXTAUTH_SECRET`:
```bash
openssl rand -base64 32
```

### Step 5: Current Implementation
The Login component currently has a placeholder Google login button. To make it fully functional:

#### Option A: Using @react-oauth/google (Recommended)
```typescript
import { GoogleLogin } from '@react-oauth/google';

// In your Login component:
<GoogleLogin
  onSuccess={credentialResponse => {
    // Handle successful login
    console.log(credentialResponse);
    onLoginAsUser();
  }}
  onError={() => {
    console.log('Login Failed');
  }}
/>
```

#### Option B: Using NextAuth.js
Create `app/api/auth/[...nextauth]/route.ts`:
```typescript
import NextAuth from "next-auth"
import GoogleProvider from "next-auth/providers/google"

const handler = NextAuth({
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID || "",
      clientSecret: process.env.GOOGLE_CLIENT_SECRET || "",
    }),
  ],
})

export { handler as GET, handler as POST }
```

Then use `signIn` from next-auth in your Login component.

### Step 6: Test the Login
1. Run your app: `npm run dev`
2. Navigate to `http://localhost:3000`
3. Click "Sign in with Google"
4. You should see the Google login popup

## Current Implementation Status

✅ Login component created with Google button placeholder
✅ Support for Admin, Tenant, and Guest logins
✅ Integrated with main app routing
✅ Logout functionality added

### Next Steps to Complete Google OAuth:
1. Create Google Cloud project and get credentials
2. Choose either @react-oauth/google or NextAuth.js
3. Implement the authentication backend
4. Add JWT token management for session handling
5. Secure API endpoints with authentication middleware

## Testing Credentials
For testing purposes, the current app accepts any email/password combination:
- Email: `test@example.com`
- Password: `anything`

This will be replaced with actual authentication once Google OAuth is configured.
