package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeanSeaLevel_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    MeanSeaLevel
		expected string
	}{
		{
			name:     "Should round to 2 decimal places",
			input:    MeanSeaLevel(1.284),
			expected: "1.28",
		},
		{
			name:     "Should round up",
			input:    MeanSeaLevel(1.286),
			expected: "1.29",
		},
		{
			name:     "Should handle exactly 2 decimal places",
			input:    MeanSeaLevel(1.28),
			expected: "1.28",
		},
		{
			name:     "Should handle repeating decimals",
			input:    MeanSeaLevel(1.3333333333),
			expected: "1.33",
		},
		{
			name:     "Should handle zero",
			input:    MeanSeaLevel(0),
			expected: "0.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(b))
		})
	}
}
