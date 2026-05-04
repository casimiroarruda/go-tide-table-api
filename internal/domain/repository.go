package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type LocationRepository interface {
	FetchAllByName(ctx context.Context, nameFilter string) (*[]Location, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Location, error)
	FindNearest(ctx context.Context, longitude float64, latitude float64) (*[]Location, error)
}

type TideRepository interface {
	GetTideTable(ctx context.Context, locationID uuid.UUID, day time.Time) ([]Tide, error)
}
