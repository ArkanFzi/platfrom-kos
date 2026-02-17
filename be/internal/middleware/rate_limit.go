package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter implements a simple in-memory rate limiter
type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	rate     int           // requests allowed
	per      time.Duration // time window
}

type visitor struct {
	lastSeen time.Time
	count    int
}

// NewRateLimiter creates a new rate limiter
// rate: number of requests allowed
// per: time window (e.g., time.Minute)
func NewRateLimiter(rate int, per time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		per:      per,
	}

	// Cleanup expired visitors every minute
	go rl.cleanupVisitors()

	return rl
}

// Middleware returns a Gin middleware function
func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !rl.allow(ip) {
			c.Header("X-RateLimit-Limit", string(rune(rl.rate)))
			c.Header("X-RateLimit-Remaining", "0")
			c.Header("Retry-After", string(rune(int(rl.per.Seconds()))))
			
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
				"message": "too many requests, please try again later",
				"retry_after_seconds": int(rl.per.Seconds()),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// allow checks if the IP is allowed to make a request
func (rl *RateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	now := time.Now()

	if !exists {
		rl.visitors[ip] = &visitor{
			lastSeen: now,
			count:    1,
		}
		return true
	}

	// Reset count if time window has passed
	if now.Sub(v.lastSeen) > rl.per {
		v.lastSeen = now
		v.count = 1
		return true
	}

	// Check if rate limit exceeded
	if v.count >= rl.rate {
		return false
	}

	// Increment count
	v.count++
	v.lastSeen = now
	return true
}

// cleanupVisitors removes old visitor records
func (rl *RateLimiter) cleanupVisitors() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > 2*rl.per {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// Predefined rate limiters for common use cases

// StrictRateLimit for sensitive endpoints (login, register)
// 5 requests per minute
func StrictRateLimit() gin.HandlerFunc {
	limiter := NewRateLimiter(5, time.Minute)
	return limiter.Middleware()
}

// ModerateRateLimit for authentication-related endpoints
// 10 requests per minute
func ModerateRateLimit() gin.HandlerFunc {
	limiter := NewRateLimiter(10, time.Minute)
	return limiter.Middleware()
}

// GeneralRateLimit for general API endpoints
// 100 requests per minute
func GeneralRateLimit() gin.HandlerFunc {
	limiter := NewRateLimiter(100, time.Minute)
	return limiter.Middleware()
}
