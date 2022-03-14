package exercise

import "sort"

type sortable []int

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func findKthLargest(nums []int, k int) int {
	sort.Sort(sortable(nums))
	return nums[len(nums)-k]
}
