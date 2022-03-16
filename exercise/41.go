package exercise

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Sort(sortable(nums))
	min := 1
	before := 0

	for _, v := range nums {
		if v > 0 {
			if v > min {
				return min
			}
			if before != v {
				min++
			}
			before = v

		}
	}
	return min
}

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

// 不懂官方解答，多看看
func firstMissingPositive_(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
