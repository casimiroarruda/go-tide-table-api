package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
	"github.com/google/uuid"

	"github.com/go-chi/chi/v5"
)

type TideHandler struct {
	tideRepo domain.TideRepository
}

func NewTideHandler(repo domain.TideRepository) *TideHandler {
	return &TideHandler{tideRepo: repo}
}

func (h *TideHandler) GetTideTable(w http.ResponseWriter, r *http.Request) {
	locationId := chi.URLParam(r, "id")
	dateStr := chi.URLParam(r, "date")

	var targetDate time.Time
	var err error

	targetDate = time.Now().UTC()
	if dateStr != "" {
		targetDate, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}
	}

	parsedUUID, err := uuid.Parse(locationId)
	if err != nil {
		http.Error(w, "Invalid location ID format", http.StatusBadRequest)
		return
	}

	tides, err := h.tideRepo.GetTideTable(
		r.Context(),
		parsedUUID,
		targetDate,
	)
	if err != nil {
		log.Printf("❌ Erro detalhado no Repo: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}

	data, err := json.Marshal(tides)
	if err != nil {
		log.Printf("❌ Erro ao codificar JSON: %v", err)
		http.Error(w, "Erro ao processar resposta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
