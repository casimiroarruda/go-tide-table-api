package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// TideHeight is a float64 that ensures 2 decimal places during JSON serialization.
type TideHeight float64

// MarshalJSON implements the json.Marshaler interface.
func (f TideHeight) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", f)), nil
}

type Location struct {
	ID           uuid.UUID  `db:"id" json:"id"`
	MarineID     string     `db:"marine_id" json:"marineId"`
	Name         string     `db:"name" json:"name"`
	Point        string     `db:"point" json:"point"` // Ex: "POINT(-23.55 -46.63)"
	MeanSeaLevel TideHeight `db:"mean_sea_level" json:"meanSeaLevel"`
	Timezone     string     `db:"timezone" json:"timezone"`
}
