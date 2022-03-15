package exercise

func deleteAndEarn(nums []int) int {
	end := 0

	for _, v := range nums {
		end = max(v, end)
	}

	vals := make([]int, end+1)

	for _, v := range nums {
		vals[v] += v
	}

	return rob(vals)
}

func rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
