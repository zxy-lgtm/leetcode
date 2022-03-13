package exercise

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			} else {
				nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
			}
		}
	}
	return 0
}
