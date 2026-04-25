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

func TestLocationRepo_FetchAll(t *testing.T) {
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
	mock.ExpectQuery(`(?i)SELECT (.+) FROM location`).WillReturnRows(rows)

	// 4. Execução
	locations, err := repo.FetchAll(context.Background(), "")

	// 5. Asserts (Validações)
	assert.NoError(t, err)
	assert.Len(t, locations, 1)
	if len(locations) > 0 {
		assert.Equal(t, "PORTO DO RECIFE", locations[0].Name)
		assert.Equal(t, domain.MeanSeaLevel(1.28), locations[0].MeanSeaLevel)
	}
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLocationRepo_FetchAll_WithFilter(t *testing.T) {
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
	mock.ExpectQuery(`(?i)SELECT .* FROM location WHERE name ILIKE .* ORDER BY name ASC`).
		WithArgs("Recife").
		WillReturnRows(rows)

	locations, err := repo.FetchAll(context.Background(), "Recife")

	assert.NoError(t, err)
	assert.Len(t, locations, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}
