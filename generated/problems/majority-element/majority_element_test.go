package problems

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Basic case",
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "Majority not at the beginning",
			nums:     []int{1, 2, 2, 3, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "Majority appears gradually",
			nums:     []int{1, 1, 2, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "Majority survives interference",
			nums:     []int{3, 3, 4, 2, 4, 4, 2, 4, 4},
			expected: 4,
		},
		{
			name:     "Single element",
			nums:     []int{7},
			expected: 7,
		},
		{
			name:     "Minimum majority (odd length)",
			nums:     []int{5, 5, 1},
			expected: 5,
		},
		{
			name:     "Minimum majority (even length)",
			nums:     []int{6, 6, 6, 1},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := majorityElement(tt.nums)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
