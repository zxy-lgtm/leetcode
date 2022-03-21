package main

import (
	"container/list"
	"fmt"
)

func canArrange(arr []int, k int) bool {
	m := make(map[int]int)

	for _, v := range arr {
		m[((v%k)+k)%k]++
	}

	if len(arr)%2 != 0 {
		return false
	}

	if m[0]%2 != 0 {
		return false
	} // 为什么要判断m[0]? 因为是2的时候 k-m = m;

	for _, v := range arr {
		mo := ((v % k) + k) % k
		if mo != 0 {
			mo = k - mo
		}
		if m[mo] <= 0 {
			return false
		}
		m[mo]--

	}

	return true
}

func trap(height []int) int {
	s := list.New()
	all := 0
	for i := 0; i < len(height)-1; i++ {
		if height[i] < height[i+1] {
			s.PushBack(height[i])
		}

	}

}

func main() {
	//arr := []int{1, 0, 2, 3, 5, 6, 0, 0, 0, 2}
	//ok := canArrange(arr[:], 2)
	num1, num2 := 1, 2
	num1, num2 = num2+1, num1+1
	num3, num4 := 1, 2
	num3 = num4 + 1
	num4 = num3 + 1
	fmt.Println(num1, num2, num3, num4)
}
