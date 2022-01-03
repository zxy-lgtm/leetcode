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
	s := "abcd"
	fmt.Println(compare_(s[0:4]))
	//fmt.Println(compare(s[0:3], "acb"))
	//fmt.Println(nums)
}

func compare_(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func compare(s string, p string) bool {
	m := make(map[rune]int)

	for _, k := range s {
		m[k]++
	}
	for _, k := range p {
		if _, ok := m[k]; !ok {
			return false
		}
		m[k]--
	}

	for _, k := range m {
		if k != 0 {
			return false
		}
	}

	return true
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
