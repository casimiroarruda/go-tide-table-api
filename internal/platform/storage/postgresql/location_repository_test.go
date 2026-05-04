package postgresql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestLocationRepo_FetchAllByName(t *testing.T) {
	// 1. Setup do Mock
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock: %v", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "postgres")
	repo := NewLocationRepo(sqlxDB)

	// 2. Dados de Exemplo (Fixtures)
	id := uuid.New()
	msl := 1.28

	// Removido "state" pois não existe na struct Location
	rows := sqlmock.NewRows([]string{"id", "marine_id", "name", "point", "mean_sea_level", "timezone"}).
		AddRow(id, "24", "PORTO DO RECIFE", "POINT(-34.87 -8.05)", msl, "-03:00")

	// 3. Expectativa: O regex agora ignora espaços extras e quebras de linha
	mock.ExpectQuery(`(?is)SELECT (.+) FROM tide_tracker.location`).WillReturnRows(rows)

	// 4. Execução
	locations, err := repo.FetchAllByName(context.Background(), "")

	// 5. Asserts (Validações)
	assert.NoError(t, err)
	assert.NotNil(t, locations)
	assert.Len(t, *locations, 1)
	if locations != nil && len(*locations) > 0 {
		assert.Equal(t, "PORTO DO RECIFE", (*locations)[0].Name)
		assert.Equal(t, domain.TideHeight(1.28), (*locations)[0].MeanSeaLevel)
	}
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLocationRepo_FetchAllByName_WithFilter(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock: %v", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "postgres")
	repo := NewLocationRepo(sqlxDB)

	id := uuid.New()
	rows := sqlmock.NewRows([]string{"id", "marine_id", "name", "point", "mean_sea_level", "timezone"}).
		AddRow(id, "24", "PORTO DO RECIFE", "POINT(-34.87 -8.05)", 1.28, "-03:00")

	// Expectativa com o WHERE clause - regex mais flexível para o ILIKE
	mock.ExpectQuery(`(?is)SELECT .* FROM tide_tracker.location WHERE name ILIKE .* ORDER BY name ASC`).
		WithArgs("%Recife%").
		WillReturnRows(rows)

	locations, err := repo.FetchAllByName(context.Background(), "Recife")

	assert.NoError(t, err)
	assert.NotNil(t, locations)
	assert.Len(t, *locations, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLocationRepo_GetByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock: %v", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "postgres")
	repo := NewLocationRepo(sqlxDB)

	id := uuid.New()
	rows := sqlmock.NewRows([]string{"id", "marine_id", "name", "point", "mean_sea_level", "timezone"}).
		AddRow(id, "24", "PORTO DO RECIFE", "POINT(-34.87 -8.05)", 1.28, "-03:00")

	mock.ExpectQuery(`(?is)SELECT .* FROM tide_tracker.location WHERE id = \$1`).
		WithArgs(id).
		WillReturnRows(rows)

	location, err := repo.GetByID(context.Background(), id)

	assert.NoError(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "PORTO DO RECIFE", location.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLocationRepo_FindNearest(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock: %v", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "postgres")
	repo := NewLocationRepo(sqlxDB)

	id := uuid.New()
	rows := sqlmock.NewRows([]string{"id", "marine_id", "name", "point", "mean_sea_level", "timezone"}).
		AddRow(id, "24", "PORTO DO RECIFE", "POINT(-34.87 -8.05)", 1.28, "-03:00")

	mock.ExpectQuery(`(?is)SELECT .* FROM tide_tracker.location ORDER BY point <-> .* LIMIT 3`).
		WithArgs(-34.87, -8.05).
		WillReturnRows(rows)

	locations, err := repo.FindNearest(context.Background(), -34.87, -8.05)

	assert.NoError(t, err)
	assert.NotNil(t, locations)
	assert.Len(t, *locations, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}
