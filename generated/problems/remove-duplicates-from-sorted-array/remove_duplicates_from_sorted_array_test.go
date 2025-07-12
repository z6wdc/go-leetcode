package problems

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
		length   int
	}{
		{[]int{1, 1, 2}, []int{1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
		{[]int{}, []int{}, 0},
		{[]int{1, 1, 1, 1, 1}, []int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5},
	}

	for _, tc := range testCases {
		numsCopy := make([]int, len(tc.nums))
		copy(numsCopy, tc.nums)
		got := removeDuplicates(numsCopy)
		if got != tc.length {
			t.Errorf("removeDuplicates(%v) = %d; want %d", tc.nums, got, tc.length)
		}
		for i := 0; i < got; i++ {
			if numsCopy[i] != tc.expected[i] {
				t.Errorf("After remove: got %v; want prefix %v", numsCopy[:got], tc.expected)
				break
			}
		}
	}
}
