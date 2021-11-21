package main

import (
	"fmt"
	"sort"
)

func main() {
	//s := "test"
	//sqt := []interface{}{s}
	//d := reflect.ValueOf(sqt)
	//fmt.Println(sqt, d.Kind())

	//bs := []byte(s)
	bs := []int{1, 0, -1, 0, -2, 2}
	//sort.Ints(bs)
	fmt.Println(fourSum(bs, 0))

}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	//fmt.Println(nums)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		fmt.Println(1)
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
				//fmt.Println(1)
				n3 := nums[l]
				n4 := nums[r]
				sum := n1 + n2 + n3 + n4
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					//fmt.Println(1)
					res = append(res, []int{n1, n2, n3, n4})
					for r > l && nums[l-1] == n3 {
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
