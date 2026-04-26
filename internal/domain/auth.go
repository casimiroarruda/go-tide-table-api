package domain

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type StringSlice []string

// Scan implements the sql.Scanner interface for PostgreSQL arrays.
func (s *StringSlice) Scan(src any) error {
	if src == nil {
		*s = nil
		return nil
	}
	asString, ok := src.(string)
	if !ok {
		return fmt.Errorf("scan error: expected string, got %T", src)
	}

	str := strings.Trim(asString, "{}")
	if str == "" {
		*s = []string{}
		return nil
	}

	parts := strings.Split(str, ",")
	for i, v := range parts {
		parts[i] = strings.Trim(v, "\"")
	}
	*s = parts
	return nil
}

type ClientCredentials struct {
	ClientID     string      `db:"client_id"`
	ClientSecret string      `db:"client_secret"`
	Name         string      `db:"name"`
	Scopes       StringSlice `db:"scopes"`
	Active       bool        `db:"active"`
	CreatedAt    time.Time   `db:"created_at"`
	UpdatedAt    time.Time   `db:"updated_at"`
}

type AuthRepository interface {
	ValidateClient(ctx context.Context, id, secret string) (*ClientCredentials, error)
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` //Bearer
	ExpiresIn   int64  `json:"expires_in"`
}
