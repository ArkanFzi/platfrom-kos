package middleware

import (
	"koskosan-be/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// ErrorHandlingMiddleware menangani semua error yang terjadi
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.GlobalLogger.Error("Panic recovered: %v", err)
				utils.InternalServerError(c, "Internal server error")
			}
		}()

		c.Next()

		// Handle errors yang di-set oleh handlers
		if len(c.Errors) > 0 {
			lastError := c.Errors.Last()
			utils.GlobalLogger.Error("Request error: %v", lastError.Error())

			// Jika status code belum di-set, default 500
			if c.Writer.Written() == false {
				utils.InternalServerError(c, lastError.Error())
			}
		}
	}
}

// LoggingMiddleware mencatat semua HTTP requests
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.RequestURI
		method := c.Request.Method
		clientIP := c.ClientIP()

		// Process request
		c.Next()

		// Log request details
		utils.GlobalLogger.LogRequest(method, path, clientIP, startTime, c.Writer.Status())
	}
}

// SecurityHeadersMiddleware menambahkan security headers
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Next()
	}
}
