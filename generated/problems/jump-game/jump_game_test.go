package problems

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		// ✅ Basic cases where it's possible to reach the last index
		{nums: []int{2, 3, 1, 1, 4}, expected: true},    // Greedy jump works
		{nums: []int{5, 0, 0, 0, 0, 0}, expected: true}, // Large jump from the beginning
		{nums: []int{2, 0, 0}, expected: true},          // Exactly reaches the end

		// ❌ Cases where it's not possible to reach the last index
		{nums: []int{3, 2, 1, 0, 4}, expected: false}, // Trapped at index 3
		{nums: []int{0, 2, 3}, expected: false},       // Cannot make any move from the start
		{nums: []int{1, 1, 0, 0}, expected: false},    // Blocked by consecutive zeros

		// 🧪 Edge cases
		{nums: []int{0}, expected: true},                          // Already at the end
		{nums: []int{1, 0}, expected: true},                       // Can move directly to the end
		{nums: []int{3, 2, 1, 0, 4, 1, 1, 1, 1}, expected: false}, // Looks long but stuck at zero
	}

	for i, tt := range tests {
		result := canJump(tt.nums)
		if result != tt.expected {
			t.Errorf("test %d: canJump(%v) = %v, expected %v", i+1, tt.nums, result, tt.expected)
		}
	}
}
