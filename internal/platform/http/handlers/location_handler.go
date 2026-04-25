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
	// Captura o filtro "name" da query string (conforme seu OpenAPI)
	nameFilter := r.URL.Query().Get("name")
	ctx := r.Context()

	// Por enquanto, chamamos o fetch geral (podemos expandir o repo para filtrar depois)
	locations, err := h.repo.FetchAll(ctx, nameFilter)
	if err != nil {
		log.Printf("❌ Erro detalhado no Repo: %v", err) // Isso vai aparecer no seu terminal
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}
