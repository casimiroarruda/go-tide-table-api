package postgresql

import (
	"context"
	"time"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TideRepo struct {
	db *sqlx.DB
}

func NewTideRepo(db *sqlx.DB) *TideRepo {
	return &TideRepo{db: db}
}

func (r *TideRepo) GetTideTable(ctx context.Context, locationID uuid.UUID, day time.Time) ([]domain.Tide, error) {
	var tides []domain.Tide

	dateStr := day.Format("2006-01-02")

	query := `
		SELECT
		t.location_id,
		t.time,
		t.height,
		CASE
			WHEN t.height >= l.mean_sea_level THEN 'HIGH'
			ELSE 'LOW'
		END as tide_type
		FROM
		tide_tracker.tide t
		JOIN tide_tracker.location l ON t.location_id = l.id
		WHERE
		t.location_id = $1
		AND t.time >= ($2::date)::timestamptz AT TIME ZONE l.timezone
		AND t.time < ($2::date + INTERVAL '1 day')::timestamptz AT TIME ZONE l.timezone
		ORDER BY
		t.time ASC
    `

	err := r.db.SelectContext(ctx, &tides, query, locationID, dateStr)
	return tides, err
}
