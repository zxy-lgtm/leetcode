package leetcode

import "sort"

func fourSum(nums []int, target int) (res [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && nums[j-1] == n2 {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				n3 := nums[l]
				n4 := nums[r]
				sum := n1 + n2 + n3 + n4
				//fmt.Println(n1, n2, n3, n4)
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					for r > l && nums[l+1] == n3 {
						l++
					}
					for r > l && nums[r-1] == n4 {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return res
}
