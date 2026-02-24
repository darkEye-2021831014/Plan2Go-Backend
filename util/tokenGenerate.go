package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims structure (token payload)
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token using only email
func GenerateToken(jwtKey string, email string) (string, error) {

	expiration := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

// VerifyToken validates the JWT token
func VerifyToken(tokenString string, jwtKey string) (*Claims, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
