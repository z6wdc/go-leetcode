package problems

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "basic example",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5, // Buy at 1, sell at 6
		},
		{
			name:   "all decreasing",
			prices: []int{7, 6, 4, 3, 1},
			want:   0, // No profit possible
		},
		{
			name:   "all increasing",
			prices: []int{1, 2, 3, 4, 5},
			want:   4, // Buy at 1, sell at 5
		},
		{
			name:   "only one day",
			prices: []int{3},
			want:   0, // Can't sell
		},
		{
			name:   "empty array",
			prices: []int{},
			want:   0,
		},
		{
			name:   "profit at the end",
			prices: []int{9, 8, 7, 1, 2},
			want:   1, // Buy at 1, sell at 2
		},
		{
			name:   "profit in the middle",
			prices: []int{2, 4, 1},
			want:   2, // Buy at 2, sell at 4
		},
		{
			name:   "max price at the start",
			prices: []int{10, 1, 2, 3, 4},
			want:   3, // Buy at 1, sell at 4
		},
		{
			name:   "multiple same prices",
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			want:   4, // Buy at 0, sell at 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit(%v) = %d; want %d", tt.prices, got, tt.want)
			}
		})
	}
}
