package exercise

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		} else if nums[mid] < nums[mid+1] && nums[mid] > nums[len(nums)-1] {
			l = mid + 1
		} else if nums[mid] < nums[mid+1] && nums[mid] < nums[len(nums)-1] {
			r = mid - 1
		}
	}
	return -1

}
