package handlers

import (
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

// MockLocationRepository is a mock implementation of domain.LocationRepository
type MockLocationRepository struct {
	mock.Mock
}

func (m *MockLocationRepository) FetchAll(ctx context.Context, nameFilter string) ([]domain.Location, error) {
	args := m.Called(ctx, nameFilter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Location), args.Error(1)
}

func (m *MockLocationRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Location, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Location), args.Error(1)
}

func TestLocationHandler_GetLocations(t *testing.T) {
	t.Run("Should return locations successfully", func(t *testing.T) {
		mockRepo := new(MockLocationRepository)
		handler := NewLocationHandler(mockRepo)

		expectedLocations := []domain.Location{
			{ID: uuid.New(), Name: "Recife", MeanSeaLevel: 1.28},
		}

		mockRepo.On("FetchAll", mock.Anything, "").Return(expectedLocations, nil)

		req, _ := http.NewRequest("GET", "/api/locations", nil)
		rr := httptest.NewRecorder()

		handler.GetLocations(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var actualLocations []domain.Location
		err := json.Unmarshal(rr.Body.Bytes(), &actualLocations)
		assert.NoError(t, err)
		assert.Equal(t, expectedLocations, actualLocations)
	})

	t.Run("Should pass name filter to repository", func(t *testing.T) {
		mockRepo := new(MockLocationRepository)
		handler := NewLocationHandler(mockRepo)

		mockRepo.On("FetchAll", mock.Anything, "Recife").Return([]domain.Location{}, nil)

		req, _ := http.NewRequest("GET", "/api/locations?name=Recife", nil)
		rr := httptest.NewRecorder()

		handler.GetLocations(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should return 500 when repository fails", func(t *testing.T) {
		mockRepo := new(MockLocationRepository)
		handler := NewLocationHandler(mockRepo)

		mockRepo.On("FetchAll", mock.Anything, "").Return(nil, errors.New("db error"))

		req, _ := http.NewRequest("GET", "/api/locations", nil)
		rr := httptest.NewRecorder()

		handler.GetLocations(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
