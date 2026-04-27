package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlice_Scan(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected StringSlice
		wantErr  bool
	}{
		{
			name:     "simple array",
			input:    "{admin,read}",
			expected: StringSlice{"admin", "read"},
			wantErr:  false,
		},
		{
			name:     "quoted array",
			input:    "{\"admin\",\"read\"}",
			expected: StringSlice{"admin", "read"},
			wantErr:  false,
		},
		{
			name:     "empty array",
			input:    "{}",
			expected: StringSlice{},
			wantErr:  false,
		},
		{
			name:     "null input",
			input:    nil,
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "invalid type",
			input:    123,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s StringSlice
			err := s.Scan(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, s)
			}
		})
	}
}
