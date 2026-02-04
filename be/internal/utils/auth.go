package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

// TokenClaims struktur untuk JWT claims
type TokenClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// SetAuthCookie set JWT token ke httpOnly cookie
func SetAuthCookie(c *gin.Context, token string, expiresIn time.Duration) {
	c.SetCookie(
		"auth_token",           // cookie name
		token,                  // cookie value
		int(expiresIn.Seconds()), // maxAge in seconds
		"/",                    // path
		"",                     // domain (empty = current domain)
		false,                  // secure (set true in production with HTTPS)
		true,                   // httpOnly - CRITICAL for security
	)
	
	// Also set expiry cookie untuk frontend reference
	expiresAt := time.Now().Add(expiresIn)
	c.SetCookie(
		"token_expires",
		expiresAt.Format(time.RFC3339),
		int(expiresIn.Seconds()),
		"/",
		"",
		false,
		false, // tidak httpOnly - frontend boleh baca
	)
}

// GetAuthToken mendapatkan token dari cookie
func GetAuthToken(c *gin.Context) (string, error) {
	token, err := c.Cookie("auth_token")
	if err != nil {
		return "", errors.New("auth token not found in cookies")
	}
	return token, nil
}

// ClearAuthCookie menghapus auth cookie (logout)
func ClearAuthCookie(c *gin.Context) {
	c.SetCookie(
		"auth_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
	c.SetCookie(
		"token_expires",
		"",
		-1,
		"/",
		"",
		false,
		false,
	)
}

// ValidateToken validate JWT token
func ValidateToken(token string, jwtSecret string) (*TokenClaims, error) {
	claims := &TokenClaims{}
	
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}
	
	// Check if token expired
	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, errors.New("token expired")
	}
	
	return claims, nil
}

// GenerateToken generate JWT token
func GenerateToken(userID int, username string, role string, jwtSecret string, expiresIn time.Duration) (string, error) {
	expiresAt := time.Now().Add(expiresIn)
	
	claims := &TokenClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}
