package problems

import "testing"

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func TestGetDecimalValue(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{1, 0, 1}, expected: 5},
		{input: []int{0}, expected: 0},
		{input: []int{1}, expected: 1},
		{input: []int{1, 1, 1, 1}, expected: 15},
		{input: []int{0, 0, 0, 0}, expected: 0},
		{input: []int{1, 0, 0, 1, 1}, expected: 19},    // 10011
		{input: []int{1, 0, 1, 0, 1, 0}, expected: 42}, // 101010
	}

	for _, tc := range tests {
		head := buildList(tc.input)
		got := getDecimalValue(head)
		if got != tc.expected {
			t.Errorf("Input: %v, Expected: %d, Got: %d", tc.input, tc.expected, got)
		}
	}
}
