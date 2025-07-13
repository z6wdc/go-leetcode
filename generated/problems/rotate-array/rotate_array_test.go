package problems

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			nums:     []int{1, 2, 3},
			k:        0,
			expected: []int{1, 2, 3},
		},
		{
			nums:     []int{1, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
		{
			nums:     []int{1, 2, 3},
			k:        4,
			expected: []int{3, 1, 2},
		},
		{
			nums:     []int{1},
			k:        10,
			expected: []int{1},
		},
		{
			nums:     []int{},
			k:        1,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		input := append([]int{}, tt.nums...)
		rotate(input, tt.k)
		if !reflect.DeepEqual(input, tt.expected) {
			t.Errorf("rotate(%v, %d) = %v; want %v", tt.nums, tt.k, input, tt.expected)
		}
	}
}
