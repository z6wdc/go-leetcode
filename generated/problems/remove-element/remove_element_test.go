package problems

import "testing"

func TestRemoveElement(t *testing.T) {
	type testCase struct {
		nums     []int
		val      int
		wantLen  int
		expected []int
	}

	tests := []testCase{
		{
			nums:     []int{3, 2, 2, 3},
			val:      3,
			wantLen:  2,
			expected: []int{2, 2},
		},
		{
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			wantLen:  5,
			expected: []int{0, 1, 3, 0, 4},
		},
		{
			nums:     []int{1, 1, 1},
			val:      1,
			wantLen:  0,
			expected: []int{},
		},
		{
			nums:     []int{},
			val:      10,
			wantLen:  0,
			expected: []int{},
		},
		{
			nums:     []int{4, 5},
			val:      3,
			wantLen:  2,
			expected: []int{4, 5},
		},
		{
			nums:     []int{1, 1, 1, 1},
			val:      1,
			wantLen:  0,
			expected: []int{},
		},
	}

	for i, tc := range tests {
		numsCopy := append([]int(nil), tc.nums...)
		gotLen := removeElement(numsCopy, tc.val)

		if gotLen != tc.wantLen {
			t.Errorf("test %d: expected length %d, got %d", i, tc.wantLen, gotLen)
		}

		gotResult := numsCopy[:gotLen]
		if !equalIgnoreOrder(gotResult, tc.expected) {
			t.Errorf("test %d: expected elements %v, got %v", i, tc.expected, gotResult)
		}
	}
}

func equalIgnoreOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	count := map[int]int{}
	for _, x := range a {
		count[x]++
	}
	for _, x := range b {
		count[x]--
		if count[x] < 0 {
			return false
		}
	}
	return true
}
