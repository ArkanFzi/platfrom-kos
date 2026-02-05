"use client";

import { useSearchParams } from 'next/navigation';
import { ResetPassword } from '@/app/components/shared/ResetPassword';
import { Suspense } from 'react';

function ResetPasswordPageContent() {
  const searchParams = useSearchParams();
  const token = searchParams.get('token');

  if (!token) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-50">
        <div className="text-center">
          <h1 className="text-2xl font-bold text-red-600">Invalid Link</h1>
          <p className="text-gray-600 mt-2">No reset token found.</p>
        </div>
      </div>
    );
  }

  return <ResetPassword token={token} />;
}

export default function ResetPasswordPage() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <ResetPasswordPageContent />
    </Suspense>
  );
}
