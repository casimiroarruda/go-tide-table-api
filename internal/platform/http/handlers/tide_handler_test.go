package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type MockTideRepository struct {
	mockGetTideTable func(ctx context.Context, locationID uuid.UUID, day time.Time) ([]domain.Tide, error)
}

func (m *MockTideRepository) GetTideTable(ctx context.Context, locationID uuid.UUID, day time.Time) ([]domain.Tide, error) {
	if m.mockGetTideTable != nil {
		return m.mockGetTideTable(ctx, locationID, day)
	}
	return nil, nil
}

func TestTideHandler_GetTideTable(t *testing.T) {
	locationID := uuid.New()
	validDateStr := "2026-05-01"

	tests := []struct {
		name           string
		locationID     string
		date           string
		mockSetup      func(m *MockTideRepository)
		expectedStatus int
		expectJSONErr  bool
	}{
		{
			name:       "Success - valid date",
			locationID: locationID.String(),
			date:       validDateStr,
			mockSetup: func(m *MockTideRepository) {
				m.mockGetTideTable = func(ctx context.Context, locID uuid.UUID, day time.Time) ([]domain.Tide, error) {
					return []domain.Tide{
						{LocationID: locID, Time: day, Type: "HIGH"},
					}, nil
				}
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:       "Success - empty date (uses current date)",
			locationID: locationID.String(),
			date:       "",
			mockSetup: func(m *MockTideRepository) {
				m.mockGetTideTable = func(ctx context.Context, locID uuid.UUID, day time.Time) ([]domain.Tide, error) {
					return []domain.Tide{}, nil
				}
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Error - invalid date format",
			locationID:     locationID.String(),
			date:           "01-05-2026",
			mockSetup:      func(m *MockTideRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Error - invalid location ID format",
			locationID:     "invalid-uuid",
			date:           validDateStr,
			mockSetup:      func(m *MockTideRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:       "Error - repository failure",
			locationID: locationID.String(),
			date:       validDateStr,
			mockSetup: func(m *MockTideRepository) {
				m.mockGetTideTable = func(ctx context.Context, locID uuid.UUID, day time.Time) ([]domain.Tide, error) {
					return nil, errors.New("database error")
				}
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockTideRepository{}
			tt.mockSetup(mockRepo)

			handler := NewTideHandler(mockRepo)

			req := httptest.NewRequest(http.MethodGet, "/location/"+tt.locationID+"/tides/"+tt.date, nil)
			
			// Setup chi context for URL params
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.locationID)
			rctx.URLParams.Add("date", tt.date)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()

			handler.GetTideTable(w, req)

			if status := w.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusOK {
				var response []domain.Tide
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Errorf("failed to decode response: %v", err)
				}
			}
		})
	}
}
