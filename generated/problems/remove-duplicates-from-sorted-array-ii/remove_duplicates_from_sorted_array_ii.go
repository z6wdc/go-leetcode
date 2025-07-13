package problems

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow := 2 // Start from the third element
	for fast := 2; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
