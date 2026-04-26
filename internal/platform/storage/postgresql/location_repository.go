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

	query := `SELECT id, marine_id, name, ST_AsText(point) as point, mean_sea_level, timezone
              FROM location
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

	// Query base
	query := `SELECT id, marine_id, name, ST_AsText(point) as point, mean_sea_level, timezone 
              FROM location`

	// Se houver filtro, adicionamos o WHERE
	if name != "" {
		query += " WHERE name ILIKE '%' || :name || '%'"
	}

	query += " ORDER BY name ASC"

	// Usando SelectContext para mapear diretamente o resultado para o slice
	var err error
	if name != "" {
		var boundQuery string
		var args []interface{}

		boundQuery, args, err = sqlx.Named(query, map[string]interface{}{"name": name})
		if err != nil {
			return nil, err
		}

		boundQuery = r.db.Rebind(boundQuery)
		err = r.db.SelectContext(ctx, &locations, boundQuery, args...)
	} else {
		err = r.db.SelectContext(ctx, &locations, query)
	}

	return locations, err
}
