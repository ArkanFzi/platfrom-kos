package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

// SaveUploadedFile saves a multipart file to the specified local folder path
// and returns the relative URL to access it (e.g., /uploads/folder/filename.jpg).
func SaveUploadedFile(fileHeader *multipart.FileHeader, folder string) (string, error) {
	// Open the uploaded file
	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	// Ensure the target directory exists
	// We'll store everything inside a "public" folder at the root of the project
	uploadDir := filepath.Join("public", folder)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Generate a unique filename using UUID and timestamp to prevent overwrites
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext == "" {
		ext = ".jpg" // default fallback
	}

	newFileName := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102150405"), uuid.New().String()[:8], ext)
	dstPath := filepath.Join(uploadDir, newFileName)

	// Create the destination file
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dst.Close()

	// Copy the file content
	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	// Return the accessible URL path
	// E.g., if folder is "uploads/rooms", it returns "/uploads/rooms/filename.jpg"
	return fmt.Sprintf("/%s/%s", folder, newFileName), nil
}

// DeleteLocalFile removes a file from the disk given its relative URL
// E.g., "/uploads/rooms/filename.jpg" -> deletes "public/uploads/rooms/filename.jpg"
func DeleteLocalFile(fileURL string) error {
	if fileURL == "" {
		return nil
	}

	// Remove the leading slash if present to safely join with "public"
	relativePath := strings.TrimPrefix(fileURL, "/")

	// Construct the full path
	fullPath := filepath.Join("public", relativePath)

	// Check if file exists before trying to delete
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil // File already doesn't exist, treat as success
	}

	return os.Remove(fullPath)
}

// IsImageFile validates if an uploaded file is an image by checking its MIME type
func IsImageFile(file *multipart.FileHeader) bool {
	contentType := file.Header.Get("Content-Type")
	return strings.HasPrefix(contentType, "image/")
}
