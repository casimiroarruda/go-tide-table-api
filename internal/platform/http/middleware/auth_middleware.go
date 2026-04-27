package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/casimiroarruda/go-tide-table-api/internal/platform/auth"
	"github.com/golang-jwt/jwt/v5"
)

func EnsureValidToken(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}
			if !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
				http.Error(w, "Authorization header must use Bearer scheme", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimSpace(authHeader[len("Bearer "):])
			if tokenStr == "" {
				http.Error(w, "Missing bearer token", http.StatusUnauthorized)
				return
			}
			claims := &auth.CustomClaims{}

			token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			hasScope := false
			requiredScope := "locations:read"
			for _, s := range claims.Scopes {
				if s == requiredScope {
					hasScope = true
					break
				}
			}
			if !hasScope {
				http.Error(w, "Insufficient permissions", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}
