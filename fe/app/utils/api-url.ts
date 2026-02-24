/**
 * Utility functions for handling API URLs and image paths
 */

// Get the API base URL from environment
export const getApiUrl = (): string => {
  return 'http://localhost:8080/api';
};

// Convert relative image URLs to absolute URLs
export const getImageUrl = (imageUrl: string | null | undefined): string => {
  if (!imageUrl) return '';
  
  // If already absolute URL, return as-is
  if (imageUrl.startsWith('http://') || imageUrl.startsWith('https://')) {
    return imageUrl;
  }
  
  // Otherwise, prepend API URL
  const apiUrl = getApiUrl().replace(/\/api$/, '');
  return `${apiUrl}${imageUrl}`;
};

// Get default placeholder image
export const getPlaceholderImage = (): string => {
  return 'https://images.unsplash.com/photo-1522771739844-6a9f6d5f14af?q=80&w=1080';
};
