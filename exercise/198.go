package exercise

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxnums := make([]int, len(nums)+1)
	maxnums[0] = nums[0]
	maxnums[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		maxnums[i] = max(maxnums[i-2]+nums[i], maxnums[i-1])
	}
	return maxnums[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
