package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
	"github.com/casimiroarruda/go-tide-table-api/internal/platform/auth"
)

type AuthHandler struct {
	repo domain.AuthRepository
}

func NewAuthHandler(repo domain.AuthRepository) *AuthHandler {
	return &AuthHandler{repo: repo}
}

func (h *AuthHandler) IssueToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	client, err := h.repo.ValidateClient(r.Context(), req.ClientID, req.ClientSecret)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	secret := os.Getenv("JWT_SECRET")
	token, err := auth.GenerateToken(client.ClientID.String(), []string(client.Scopes), secret)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": token,
		"token_type":   "Bearer",
	})
}
