package second

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindBestSeatDist(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "middle case 01",
			input:    []int{1, 0, 0, 0, 1},
			expected: 2,
		},
		{
			name:     "middle case 02",
			input:    []int{1, 0, 1, 0, 0, 1, 0, 0, 0, 1},
			expected: 2,
		},
		{
			name:     "middle case 03",
			input:    []int{1, 0, 1, 0},
			expected: 1,
		},

		{
			name:     "border case 01",
			input:    []int{1, 1, 0, 0, 0},
			expected: 3,
		},
		{
			name:     "border case 02",
			input:    []int{0, 0, 0, 1, 1},
			expected: 3,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := findBestSeatDist(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
