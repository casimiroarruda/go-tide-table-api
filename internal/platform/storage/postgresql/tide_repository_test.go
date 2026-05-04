package postgresql

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestTideRepo_GetTideTable(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock: %v", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "postgres")
	repo := NewTideRepo(sqlxDB)

	id := uuid.New()
	date := time.Now()
	
	rows := sqlmock.NewRows([]string{"location_id", "time", "height", "tide_type"}).
		AddRow(id, date, 1.28, "HIGH")

	mock.ExpectQuery(`(?is)SELECT .* FROM tide_tracker.tide .*`).
		WithArgs(id, date.Format("2006-01-02")).
		WillReturnRows(rows)

	tides, err := repo.GetTideTable(context.Background(), id, date)

	assert.NoError(t, err)
	assert.Len(t, tides, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}
