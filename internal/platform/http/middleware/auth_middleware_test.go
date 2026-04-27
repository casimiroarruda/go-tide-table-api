package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/casimiroarruda/go-tide-table-api/internal/platform/auth"
	"github.com/stretchr/testify/assert"
)

func TestEnsureValidToken(t *testing.T) {
	secret := "test-secret"
	middleware := EnsureValidToken(secret)
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	t.Run("valid token", func(t *testing.T) {
		token, _ := auth.GenerateToken("client-1", []string{"read"}, secret)
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rr := httptest.NewRecorder()

		middleware(nextHandler).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("missing header", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		middleware(nextHandler).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Contains(t, rr.Body.String(), "Missing Authorization header")
	})

	t.Run("invalid token format", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")
		rr := httptest.NewRecorder()

		middleware(nextHandler).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Contains(t, rr.Body.String(), "Invalid token")
	})

	t.Run("wrong secret", func(t *testing.T) {
		token, _ := auth.GenerateToken("client-1", []string{"read"}, "wrong-secret")
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rr := httptest.NewRecorder()

		middleware(nextHandler).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}
