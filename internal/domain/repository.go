package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type LocationRepository interface {
	FetchAll(ctx context.Context, nameFilter string) ([]Location, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Location, error)
}

type TideRepository interface {
	GetTideTable(ctx context.Context, locationID uuid.UUID, day time.Time) ([]Tide, error)
}
