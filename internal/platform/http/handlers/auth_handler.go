package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
	"github.com/casimiroarruda/go-tide-table-api/internal/platform/auth"
)

type AuthHandler struct {
	repo      domain.AuthRepository
	jwtSecret string
}

func NewAuthHandler(repo domain.AuthRepository, jwtSecret string) *AuthHandler {
	return &AuthHandler{repo: repo, jwtSecret: jwtSecret}
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

	token, err := auth.GenerateToken(client.ClientID.String(), []string(client.Scopes), h.jwtSecret)
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
