"use client";

import { GoogleOAuthProvider } from '@react-oauth/google';

export function GoogleAuthProvider({ children }: { children: React.ReactNode }) {
    const clientId = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || "";

    if (!clientId) {
        console.error("GOOGLE_CLIENT_ID is missing in environment variables! Add NEXT_PUBLIC_GOOGLE_CLIENT_ID to your .env.local");
    }

    return (
        <GoogleOAuthProvider clientId={clientId}>
            {children}
        </GoogleOAuthProvider>
    );
}
