package exercise

func jump(nums []int) int {
	m := 0
	next := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		if i <= m {
			m = max(m, i+nums[i])
		}
		if i >= next {
			next = m
			step++
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
