package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Scopes []string `json:"scopes"`
	jwt.RegisteredClaims
}

func GenerateToken(clientID string, scopes []string, secret string) (string, error) {
	claims := CustomClaims{
		Scopes: scopes,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "tide-table-api",
			Subject:   clientID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
