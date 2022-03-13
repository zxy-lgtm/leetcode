package exercise

func findDuplicate(nums []int) int {
	fast := nums[0]
	slow := nums[0]

	for fast < len(nums) {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			fast = nums[0]
			for slow != fast {
				slow = nums[slow]
				fast = nums[fast]
			}
			return fast
		}
	}
	return -1

}
