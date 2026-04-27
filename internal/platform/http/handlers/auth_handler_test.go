package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) ValidateClient(ctx context.Context, id, secret string) (*domain.ClientCredentials, error) {
	args := m.Called(ctx, id, secret)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ClientCredentials), args.Error(1)
}

func TestAuthHandler_IssueToken(t *testing.T) {
	repo := new(MockAuthRepository)
	handler := NewAuthHandler(repo, "test-secret")

	t.Run("successful token issue", func(t *testing.T) {
		clientID := uuid.New()
		client := &domain.ClientCredentials{
			ClientID: clientID,
			Scopes:   domain.StringSlice{"read"},
		}

		repo.On("ValidateClient", mock.Anything, clientID.String(), "secret-1").Return(client, nil)

		body := map[string]string{
			"client_id":     clientID.String(),
			"client_secret": "secret-1",
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/auth/token", bytes.NewBuffer(jsonBody))
		rr := httptest.NewRecorder()

		handler.IssueToken(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var resp map[string]string
		json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["access_token"])
		assert.Equal(t, "Bearer", resp["token_type"])
	})

	t.Run("unauthorized", func(t *testing.T) {
		wrongID := uuid.New().String()
		repo.On("ValidateClient", mock.Anything, wrongID, "wrong").Return(nil, errors.New("unauthorized"))

		body := map[string]string{
			"client_id":     wrongID,
			"client_secret": "wrong",
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/auth/token", bytes.NewBuffer(jsonBody))
		rr := httptest.NewRecorder()

		handler.IssueToken(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("invalid json", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/auth/token", bytes.NewBuffer([]byte("invalid")))
		rr := httptest.NewRecorder()

		handler.IssueToken(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
