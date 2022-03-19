package leetcode

func maxSubArray(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > sum {
			sum = nums[i]
		}
	}
	return sum
}

func maxSubArray_(nums []int) int {
	maxnum := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1]+nums[i], nums[i])
		if maxnum < nums[i] {
			maxnum = nums[i]
		}
	}
	return maxnum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
