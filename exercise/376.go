package exercise

func wiggleMaxLength(nums []int) int {
	up := 1
	down := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			up = down + 1
		}

		if nums[i-1] > nums[i] {
			down = up + 1
		}
	}
	return max(up, down)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
