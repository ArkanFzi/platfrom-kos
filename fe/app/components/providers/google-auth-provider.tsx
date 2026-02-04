"use client";

import { GoogleOAuthProvider } from '@react-oauth/google';

export function GoogleAuthProvider({ children }: { children: React.ReactNode }) {
    // Replace this with your actual Google Client ID from Google Cloud Console
    const clientId = "YOUR_GOOGLE_CLIENT_ID.apps.googleusercontent.com";

    return (
        <GoogleOAuthProvider clientId={clientId}>
            {children}
        </GoogleOAuthProvider>
    );
}
