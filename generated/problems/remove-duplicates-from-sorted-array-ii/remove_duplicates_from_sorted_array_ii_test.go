package problems

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
		want     []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7, []int{0, 0, 1, 1, 2, 3, 3}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1}, 2, []int{1, 1}},
		{[]int{1, 1, 1, 1, 1}, 2, []int{1, 1}},
		{[]int{}, 0, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 2, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3}},
		{[]int{1, 1, 1, 2, 3, 3, 3, 3, 4, 4}, 7, []int{1, 1, 2, 3, 3, 4, 4}},
	}

	for _, tt := range tests {
		numsCopy := make([]int, len(tt.nums))
		copy(numsCopy, tt.nums)
		k := removeDuplicates(numsCopy)
		if k != tt.expected || !slices.Equal(numsCopy[:k], tt.want) {
			t.Errorf("removeDuplicates(%v) = %v, %v; want %v, %v",
				tt.nums, k, numsCopy[:k], tt.expected, tt.want)
		}
	}
}
