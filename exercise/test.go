package main

import (
	"fmt"
	"math/rand"
)

func partition(nums []int, l, r int) int {
	i := l - 1
	pivot := nums[r]

	for j := l; j < r; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]

	return i + 1
}

func randomPartition(nums []int, l, r int) int {
	i := rand.Intn(r-l+1) + l

	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}

func qsort(nums []int, l, r int) {
	if l < r {
		position := randomPartition(nums, l, r)
		qsort(nums, l, position-1)
		qsort(nums, position+1, r)
	}
}

func sortArray(nums []int) []int {
	qsort(nums, 0, len(nums)-1)

	return nums
}

func partition(nums []int, l, r int) int {
	pivot := nums[r]

	i := l - 1

	for j := l; j < r; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]

	return i + 1
}

func randomPartition(nums []int, l, r int) int {
	i := rand.Intn(r-l+1) + l

	nums[i], nums[r] = nums[r], nums[i]

	return partition(nums, l, r)
}

func quickSort(nums []int, l, r int) {
	if l < r {
		position := partition(nums, l, r)
		quickSort(nums, l, position-1)
		quickSort(nums, position+1, r)
	}
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)

	return nums
}

//有重复的全排序

func Swap(list []int, i, j int) {
	list[i], list[j] = list[j], list[i]
}

func isSwap(list []int, start, end int) bool {
	for i := start; i < end; i++ {
		if list[i] == list[end] {
			return false
		}
	}
	return true
}

func AllRange(list []int, start, end int) {
	if start == end {
		fmt.Println(list)
	} else {
		for i := start; i <= end; i++ {
			if isSwap(list, start, i) {
				Swap(list, start, i)
				AllRange(list, start+1, end)
				Swap(list, start, i)
			}
		}
	}
}

func main() {
	//list := [6]int{1, 1, 1, 0, 0, 0}
	//AllRange(list[:], 0, len(list)-1)

}
