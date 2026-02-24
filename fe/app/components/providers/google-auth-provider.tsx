"use client";

import { GoogleOAuthProvider } from '@react-oauth/google';

export function GoogleAuthProvider({ children }: { children: React.ReactNode }) {
    const clientId = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || "";

    if (!clientId && process.env.NODE_ENV === 'production') {
        console.error("GOOGLE_CLIENT_ID is missing in environment variables!");
    }

    return (
        <GoogleOAuthProvider clientId={clientId}>
            {children}
        </GoogleOAuthProvider>
    );
}
