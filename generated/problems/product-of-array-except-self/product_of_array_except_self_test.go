package problems

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "basic case",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "with zero",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "single element",
			nums:     []int{42},
			expected: []int{1}, // edge case: no other elements
		},
		{
			name:     "two elements",
			nums:     []int{2, 3},
			expected: []int{3, 2},
		},
		{
			name:     "multiple zeros",
			nums:     []int{0, 4, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3},
			expected: []int{6, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("input = %v, expected = %v, got = %v", tt.nums, tt.expected, result)
			}
		})
	}
}
