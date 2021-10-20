package leetcode

func minMoves(nums []int) int {
	min := nums[0]
	max := nums[0]
	for _, value := range nums {
		if min > value {
			min = value
		}

		if max < value {
			max = value
		}
	}

	if max == min {
		return 0
	}

	total := 0

	for _, value := range nums {
		if value > min {
			total += value - min
		}
	}

	return total
}
