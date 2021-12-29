package exercise

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement_moore(nums []int) int {
	tip := [2]int{0, 0}

	for _, k := range nums {
		if tip[1] > 0 {
			if tip[0] != k {
				tip[1]--
			} else {
				tip[1]++
			}
		}

		if tip[1] == 0 {
			tip[0] = k
			tip[1]++
		}
	}

	return tip[0]
}
