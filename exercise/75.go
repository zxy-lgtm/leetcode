package exercise

func sortColors(nums []int) {
	head, tail := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[head] = nums[head], nums[i]
			head++
		}
	}
	tmp := head
	for tmp <= tail {
		if nums[tmp] == 1 {
			nums[head], nums[tmp] = nums[tmp], nums[head]
			head++
		}
		tmp++
	}
}
