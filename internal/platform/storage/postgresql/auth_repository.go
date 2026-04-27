package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/casimiroarruda/go-tide-table-api/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) ValidateClient(ctx context.Context, id, secret string) (*domain.ClientCredentials, error) {
	var client domain.ClientCredentials

	query := `SELECT client_id, client_secret, name, scopes FROM auth_store.clients WHERE client_id = $1`
	err := r.db.GetContext(ctx, &client, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("client not found")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.ClientSecret), []byte(secret))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &client, nil
}
