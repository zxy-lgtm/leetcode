package leetcode

func searchInsert(nums []int, target int) int {
	tail := len(nums)
	head := 0
	for head <= tail {
		mid := (tail + head) / 2
		if mid == len(nums) {
			return mid
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}
	return head
}
