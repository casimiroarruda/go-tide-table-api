package postgresql

import (
	"context"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LocationRepo struct {
	db *sqlx.DB
}

// GetByID implements [domain.LocationRepository].
func (r *LocationRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Location, error) {
	var location domain.Location
	query := `SELECT id, marine_id, name, tide_tracker.ST_AsText(point::tide_tracker.geography) as point, mean_sea_level, timezone
              FROM tide_tracker.location
              WHERE id = $1`

	if err := r.db.GetContext(ctx, &location, query, id); err != nil {
		return nil, err
	}

	return &location, nil
}

func NewLocationRepo(db *sqlx.DB) *LocationRepo {
	return &LocationRepo{db: db}
}

func (r *LocationRepo) FetchAll(ctx context.Context, name string) ([]domain.Location, error) {
	var locations []domain.Location

	query := `SELECT id, marine_id, name, 
				tide_tracker.ST_AsText(point::tide_tracker.geography) as point, 
				mean_sea_level, timezone 
              FROM tide_tracker.location`

	if name != "" {
		query += " WHERE name ILIKE $1"
	}

	query += " ORDER BY name ASC"

	var err error
	if name == "" {
		err = r.db.SelectContext(ctx, &locations, query)
		return locations, err
	}
	err = r.db.SelectContext(ctx, &locations, query, "%"+name+"%")
	return locations, err

}
