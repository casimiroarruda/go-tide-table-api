package domain

import (
	"context"

	"github.com/google/uuid"
)

type LocationRepository interface {
	FetchAll(ctx context.Context, nameFilter string) ([]Location, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Location, error)
}
