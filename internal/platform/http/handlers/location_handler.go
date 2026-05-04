package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
)

type LocationHandler struct {
	repo domain.LocationRepository
}

func NewLocationHandler(repo domain.LocationRepository) *LocationHandler {
	return &LocationHandler{repo: repo}
}

func (h *LocationHandler) GetLocationsByName(ctx context.Context, nameFilter string) (*[]domain.Location, error) {

	return h.repo.FetchAllByName(ctx, nameFilter)
}

func (h *LocationHandler) GetNearestLocations(ctx context.Context, latitude string, longitude string) (*[]domain.Location, error) {
	lat, _ := strconv.ParseFloat(latitude, 64)
	lon, _ := strconv.ParseFloat(longitude, 64)

	return h.repo.FindNearest(ctx, lon, lat)

}

func (h *LocationHandler) SearchByNameOrByPosition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var locations *[]domain.Location
	var err error

	if r.URL.Query().Has("name") {
		locations, err = h.GetLocationsByName(ctx, r.URL.Query().Get("name"))
	}

	if locations == nil && r.URL.Query().Has("lat") && r.URL.Query().Has("lon") {
		locations, err = h.GetNearestLocations(ctx, r.URL.Query().Get("lat"), r.URL.Query().Get("lon"))
	}

	if err != nil {
		log.Printf("❌ Erro detalhado no Repo: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}

	if locations == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)
}
