package leetcode

func removeDuplicates(nums []int) int {

	l := len(nums)

	if l == 0 {
		return 0
	}

	rl := 1
	pre := nums[0]

	for i := 1; i < l; i++ {
		if nums[i] != pre {
			nums[rl] = nums[i]
			pre = nums[i]
			rl++
		}
	}

	return rl

}
