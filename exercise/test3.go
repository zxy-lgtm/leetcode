package main

import (
	"fmt"
	"sort"
)

type sortable [][]int

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	nums := [][]int{{2, 4}, {1, 4}}

	sort.Sort(sortable(nums))
	fmt.Println(nums)
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
