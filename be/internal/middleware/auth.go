package middleware

import (
	"koskosan-be/internal/utils"
	"os"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verify JWT token dari cookie
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token dari cookie
		token, err := utils.GetAuthToken(c)
		if err != nil {
			utils.UnauthorizedError(c, "Missing authentication token")
			c.Abort()
			return
		}

		// Validate token
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			jwtSecret = "your-secret-key" // fallback untuk development
		}

		claims, err := utils.ValidateToken(token, jwtSecret)
		if err != nil {
			utils.UnauthorizedError(c, "Invalid or expired token")
			c.Abort()
			return
		}

		// Set claims ke context untuk digunakan di handlers
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// RoleMiddleware check user role
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			utils.ForbiddenError(c, "User role not found")
			c.Abort()
			return
		}

		role := userRole.(string)
		
		// Check if role is allowed
		roleAllowed := false
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			utils.ForbiddenError(c, "You don't have permission to access this resource")
			c.Abort()
			return
		}

		c.Next()
	}
}

// OptionalAuthMiddleware verify token tapi tidak mandatory
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.GetAuthToken(c)
		if err != nil {
			// Token tidak ada, tapi allowed - skip validation
			c.Next()
			return
		}

		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			jwtSecret = "your-secret-key"
		}

		claims, err := utils.ValidateToken(token, jwtSecret)
		if err != nil {
			// Token invalid, tapi allowed - skip validation
			c.Next()
			return
		}

		// Set claims ke context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
