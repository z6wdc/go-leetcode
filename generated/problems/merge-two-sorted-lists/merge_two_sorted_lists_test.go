package problems

import (
	"reflect"
	"testing"
)

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "BothEmpty",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "OneEmpty",
			list1:    []int{1, 2, 3},
			list2:    []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "BothNonEmpty",
			list1:    []int{1, 3, 5},
			list2:    []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "WithDuplicates",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "AllEqual",
			list1:    []int{2, 2},
			list2:    []int{2, 2},
			expected: []int{2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := buildList(tt.list1)
			l2 := buildList(tt.list2)
			merged := mergeTwoLists(l1, l2)
			result := listToSlice(merged)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
