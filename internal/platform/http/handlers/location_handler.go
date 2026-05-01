package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
)

type LocationHandler struct {
	repo domain.LocationRepository
}

func NewLocationHandler(repo domain.LocationRepository) *LocationHandler {
	return &LocationHandler{repo: repo}
}

func (h *LocationHandler) GetLocations(w http.ResponseWriter, r *http.Request) {
	nameFilter := r.URL.Query().Get("name")
	ctx := r.Context()

	locations, err := h.repo.FetchAll(ctx, nameFilter)
	if err != nil {
		log.Printf("❌ Erro detalhado no Repo: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)
}
