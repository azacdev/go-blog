package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/azacdev/go-blog/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

// Define your custom claims
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

var (
	jwtAccessSecret  []byte
	jwtRefreshSecret []byte
)

func init() {
	config.Set()

	cfg := config.Get()

	jwtAccessSecret = []byte(cfg.JWTAccess.Secret)
	jwtRefreshSecret = []byte(cfg.JWTRefresh.Secret)

	if len(jwtAccessSecret) == 0 {
		panic("JWT_ACCESS_SECRET environment variable not set")
	}
	if len(jwtRefreshSecret) == 0 {
		panic("JWT_REFRESH_SECRET environment variable not set")
	}
}

// GenerateTokens creates a new access token and refresh token for the given user ID.
func GenerateTokens(userID uint) (string, string, error) {
	// Access Token (short-lived)
	accessTokenClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)), // Access token expires in 15 minutes
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprintf("%d", userID),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	signedAccessToken, err := accessToken.SignedString(jwtAccessSecret)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	// Refresh Token (long-lived)
	refreshTokenClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // Refresh token expires in 7 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprintf("%d", userID),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	signedRefreshToken, err := refreshToken.SignedString(jwtRefreshSecret)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return signedAccessToken, signedRefreshToken, nil
}

// ValidateAccessToken parses and validates an access token.
func ValidateAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtAccessSecret, nil
	})

	if err != nil {
		// Specific error handling for common JWT errors
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("access token has expired")
		}
		if errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, errors.New("invalid access token format or signature")
		}
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid access token claims or token is not valid")
	}

	return claims, nil
}

// ValidateRefreshToken parses and validates a refresh token.
func ValidateRefreshToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtRefreshSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("refresh token has expired")
		}
		if errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, errors.New("invalid refresh token format or signature")
		}
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token claims or token is not valid")
	}

	return claims, nil
}
