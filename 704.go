package main

import "fmt"

func search(nums []int, target int) int {
	tail := len(nums) - 1
	head := 0

	for {
		fmt.Println(head, tail)
		if nums[tail] > target {
			head = tail
			tail = tail / 2
		}
		if nums[tail] < target {
			tail++
		}
		//可能越界
		if tail > len(nums)-1 {
			return -1
		}
		if nums[tail] == target {
			break
		}
		if head <= tail {
			return -1
		}
		//数组长度为1
		if tail == 0 && nums[0] != target {
			return -1
		}

	}
	return tail
}

/*func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	n := search(nums, 2)
	fmt.Println(n)
}*/
