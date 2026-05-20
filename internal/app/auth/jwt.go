package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("super_secret_key_32_chars_long!!") // минимум 32 символа для HS256

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT создаёт токен на 24 часа
func GenerateJWT(userID uint, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT проверяет токен и возвращает claims
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// GetUserIDFromCookie извлекает userID из cookie "token"
func GetUserIDFromCookie(cookies []*http.Cookie) (uint, error) {
	for _, c := range cookies {
		if c.Name == "token" {
			claims, err := ValidateJWT(c.Value)
			if err != nil {
				return 0, err
			}
			return claims.UserID, nil
		}
	}
	return 0, errors.New("token cookie not found")
}