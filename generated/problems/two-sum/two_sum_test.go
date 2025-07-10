package problems

import "testing"

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "basic",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "multiple choices",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "duplicate values",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "zero sum",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "no solution",
			nums:     []int{1, 2, 3},
			target:   10,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)

			if !samePair(got, tc.expected) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.expected)
			}
		})
	}
}

func samePair(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) != 2 || len(b) != 2 {
		return false
	}
	return (a[0] == b[0] && a[1] == b[1]) || (a[0] == b[1] && a[1] == b[0])
}
