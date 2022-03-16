package exercise

func canJump(nums []int) bool {
	m := 0
	for i := 0; i < len(nums); i++ {
		if i <= m {
			m = max(m, i+nums[i])
		}
		if m > len(nums)-2 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
