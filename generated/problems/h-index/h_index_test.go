package problems

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		expected  int
	}{
		{
			name:      "Example case",
			citations: []int{3, 0, 6, 1, 5},
			expected:  3,
		},
		{
			name:      "All zeros",
			citations: []int{0, 0, 0, 0},
			expected:  0,
		},
		{
			name:      "All high citations",
			citations: []int{10, 8, 9, 7},
			expected:  4,
		},
		{
			name:      "Single paper high",
			citations: []int{10},
			expected:  1,
		},
		{
			name:      "Single paper zero",
			citations: []int{0},
			expected:  0,
		},
		{
			name:      "Uniform citation",
			citations: []int{2, 2, 2},
			expected:  2,
		},
		{
			name:      "Edge case mixed",
			citations: []int{4, 4, 0, 0},
			expected:  2,
		},
		{
			name:      "Very large citations",
			citations: []int{100, 200, 300},
			expected:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hIndex(tt.citations)
			if result != tt.expected {
				t.Errorf("hIndex(%v) = %d; expected %d", tt.citations, result, tt.expected)
			}
		})
	}
}
