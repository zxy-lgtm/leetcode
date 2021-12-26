package main

import "fmt"

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
	list := [6]int{1, 1, 1, 0, 0, 0}
	AllRange(list[:], 0, len(list)-1)
}
