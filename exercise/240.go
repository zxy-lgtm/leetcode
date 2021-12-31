package exercise

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix[0])
	loc := len(matrix)
	i, j, ok := 0, 0, 0
	for i < row || j < loc {
		if i < row {
			ok = searchrow(matrix, target, i, j-1)
			i++
			if ok != -1 {
				return true
			}
		}
		if j < loc {
			ok = searchl(matrix, target, j, i-1)
			j++
			if ok != -1 {
				return true
			}
		}
		if i < row && j < loc && matrix[i][0] > target && matrix[0][j] > target {
			break
		}

	}
	return false
}

func searchrow(matrix [][]int, target int, loc int, n int) int {
	for i := 0; i <= n; i++ {
		if matrix[i][loc] == target {
			return i
		}
	}
	return -1
}

func searchl(matrix [][]int, target int, loc int, n int) int {
	for i := 0; i <= n; i++ {
		if matrix[loc][i] == target {
			return i
		}
	}
	return -1
}

// 二分查找
func Search(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		ok := sort.SearchInts(nums, target)
		if ok < len(nums) && nums[ok] == target {
			return true
		}
	}
	return false
}

// z字查找
func Search_z(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	x, y := 0, m-1
	for x < n && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
