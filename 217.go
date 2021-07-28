package leetcode

import "sort"

//暴力法
func containsDuplicate1(nums []int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

//先排序
func containsDuplicate2(nums []int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}

	sort.Ints(nums)

	for i := 0; i < l-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

//map
func containsDuplicate3(nums []int) bool {
	val := make(map[int]bool)
	for _, num := range nums {
		if val[num] {
			return true
		}
		val[num] = true
	}
	return false
}
