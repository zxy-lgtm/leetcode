package exercise

func searchMatrix(matrix [][]int, target int) bool {
	nums := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		nums = append(nums, matrix[i][0])
	}
	index := getIndex(nums, target)
	if index == -1 {
		return false
	}

	l, r := 0, len(matrix[0])
	for l <= r {
		mid := (l + r) / 2
		if mid == len(matrix[0]) {
			return false
		} else if matrix[index][mid] == target {
			return true
		} else if matrix[index][mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return matrix[index][l] == target
}

func getIndex(nums []int, target int) int {
	tail := len(nums)
	head := 0
	if nums[tail-1] < target {
		return tail - 1
	}
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
	return head - 1
}
