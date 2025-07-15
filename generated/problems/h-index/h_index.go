package problems

func hIndex(citations []int) int {
	n := len(citations)
	count := make([]int, n+1)

	for _, c := range citations {
		if c >= n {
			count[n]++
		} else {
			count[c]++
		}
	}

	h := 0
	for i := n; i >= 0; i-- {
		h += count[i]
		if h >= i {
			return i
		}
	}
	return 0
}
