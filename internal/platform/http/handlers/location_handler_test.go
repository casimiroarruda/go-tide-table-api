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

func (m *MockLocationRepository) FetchAllByName(ctx context.Context, nameFilter string) (*[]domain.Location, error) {
	args := m.Called(ctx, nameFilter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]domain.Location), args.Error(1)
}

func (m *MockLocationRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Location, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Location), args.Error(1)
}

func (m *MockLocationRepository) FindNearest(ctx context.Context, longitude float64, latitude float64) (*[]domain.Location, error) {
	args := m.Called(ctx, longitude, latitude)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]domain.Location), args.Error(1)
}

func TestLocationHandler_GetLocations(t *testing.T) {
	t.Run("Should return locations successfully", func(t *testing.T) {
		mockRepo := new(MockLocationRepository)
		handler := NewLocationHandler(mockRepo)

		expectedLocations := &[]domain.Location{
			{ID: uuid.New(), Name: "Recife", MeanSeaLevel: 1.28},
		}

		mockRepo.On("FetchAllByName", mock.Anything, "").Return(expectedLocations, nil)

		req, _ := http.NewRequest("GET", "/api/locations?name=", nil)
		rr := httptest.NewRecorder()

		handler.SearchByNameOrByPosition(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var actualLocations []domain.Location
		err := json.Unmarshal(rr.Body.Bytes(), &actualLocations)
		assert.NoError(t, err)
		assert.Equal(t, *expectedLocations, actualLocations)
	})

	t.Run("Should pass name filter to repository", func(t *testing.T) {
		mockRepo := new(MockLocationRepository)
		handler := NewLocationHandler(mockRepo)

		mockRepo.On("FetchAllByName", mock.Anything, "Recife").Return(&[]domain.Location{}, nil)

		req, _ := http.NewRequest("GET", "/api/locations?name=Recife", nil)
		rr := httptest.NewRecorder()

		handler.SearchByNameOrByPosition(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should return 500 when repository fails", func(t *testing.T) {
		mockRepo := new(MockLocationRepository)
		handler := NewLocationHandler(mockRepo)

		mockRepo.On("FetchAllByName", mock.Anything, "").Return(nil, errors.New("db error"))

		req, _ := http.NewRequest("GET", "/api/locations?name=", nil)
		rr := httptest.NewRecorder()

		handler.SearchByNameOrByPosition(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("Should return locations by position successfully", func(t *testing.T) {
		mockRepo := new(MockLocationRepository)
		handler := NewLocationHandler(mockRepo)

		expectedLocations := &[]domain.Location{
			{ID: uuid.New(), Name: "Recife", MeanSeaLevel: 1.28},
		}

		mockRepo.On("FindNearest", mock.Anything, -34.87, -8.05).Return(expectedLocations, nil)

		req, _ := http.NewRequest("GET", "/api/locations?lat=-8.05&lon=-34.87", nil)
		rr := httptest.NewRecorder()

		handler.SearchByNameOrByPosition(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var actualLocations []domain.Location
		err := json.Unmarshal(rr.Body.Bytes(), &actualLocations)
		assert.NoError(t, err)
		assert.Equal(t, *expectedLocations, actualLocations)
	})

	t.Run("Should return 422 when no parameter is provided", func(t *testing.T) {
		mockRepo := new(MockLocationRepository)
		handler := NewLocationHandler(mockRepo)

		req, _ := http.NewRequest("GET", "/api/locations", nil)
		rr := httptest.NewRecorder()

		handler.SearchByNameOrByPosition(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	})
}
