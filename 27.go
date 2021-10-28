package leetcode

func removeElement(nums []int, val int) int {
	fast, late := 0, 0
	l := len(nums)
	for fast < len(nums) {
		nums[late] = nums[fast]
		if nums[fast] == val {
			fast++
			l--
		} else {
			fast++
			late++
		}
	}
	return l
}
