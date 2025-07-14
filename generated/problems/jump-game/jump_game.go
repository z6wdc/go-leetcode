package problems

func canJump(nums []int) bool {
	maxReach := 0
	for i, n := range nums {
		if i > maxReach {
			return false
		}
		if i+n > maxReach {
			maxReach = i + n
		}
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}
