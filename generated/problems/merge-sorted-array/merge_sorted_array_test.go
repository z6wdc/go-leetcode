package problems

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
		{
			nums1: []int{2, 0},
			m:     1,
			nums2: []int{1},
			n:     1,
			want:  []int{1, 2},
		},
		{
			nums1: []int{4, 5, 6, 0, 0, 0},
			m:     3,
			nums2: []int{1, 2, 3},
			n:     3,
			want:  []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		nums1Copy := make([]int, len(tt.nums1))
		copy(nums1Copy, tt.nums1)

		merge(nums1Copy, tt.m, tt.nums2, tt.n)

		if !reflect.DeepEqual(nums1Copy, tt.want) {
			t.Errorf("merge(%v, %d, %v, %d) = %v; want %v",
				tt.nums1, tt.m, tt.nums2, tt.n, nums1Copy, tt.want)
		}
	}
}
