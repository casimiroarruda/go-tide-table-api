package postgresql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthRepo_ValidateClient(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock: %v", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "postgres")
	repo := NewAuthRepository(sqlxDB)

	clientID := uuid.New()
	password := "secret123"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	t.Run("successful validation", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"client_id", "client_secret", "name", "scopes"}).
			AddRow(clientID, string(hash), "Test Client", "{admin}")

		mock.ExpectQuery(`SELECT client_id, client_secret, name, scopes FROM auth_store.clients WHERE client_id = \$1`).
			WithArgs(clientID.String()).
			WillReturnRows(rows)

		client, err := repo.ValidateClient(context.Background(), clientID.String(), password)

		assert.NoError(t, err)
		if assert.NotNil(t, client) {
			assert.Equal(t, clientID, client.ClientID)
			assert.Equal(t, "Test Client", client.Name)
		}
	})

	t.Run("client not found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT .* FROM auth_store.clients`).
			WithArgs("unknown").
			WillReturnError(sql.ErrNoRows)

		client, err := repo.ValidateClient(context.Background(), "unknown", password)

		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "client not found")
	})

	t.Run("invalid password", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"client_id", "client_secret", "name", "scopes"}).
			AddRow(clientID, string(hash), "Test Client", "{admin}")

		mock.ExpectQuery(`SELECT .* FROM auth_store.clients`).
			WithArgs(clientID.String()).
			WillReturnRows(rows)

		client, err := repo.ValidateClient(context.Background(), clientID.String(), "wrongpassword")

		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "invalid credentials")
	})
}
