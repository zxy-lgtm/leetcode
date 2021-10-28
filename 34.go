package leetcode

func searchRange(nums []int, target int) (res []int) {
	res = append(res, []int{-1, -1}...)

	if len(nums) == 0 {
		return
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] == target {
			f, end := mid, mid
			for f-1 >= 0 && nums[f-1] == target {
				f--
			}
			for end+1 < len(nums) && nums[end+1] == target {
				end++
			}
			res[0] = f
			res[1] = end
			return res
		}
	}
	return res
}
