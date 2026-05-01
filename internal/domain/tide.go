package domain

import (
	"time"

	"github.com/google/uuid"
)

type Tide struct {
	LocationID uuid.UUID  `db:"location_id" json:"locationId"`
	Time       time.Time  `db:"time" json:"time"`
	Height     TideHeight `db:"height" json:"height"`
	Type       string     `db:"tide_type" json:"type"`
}
