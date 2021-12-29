package main

import "fmt"

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

func main() {
	arr := []int{1, 0, 2, 3, 5, 6, 0, 0, 0, 2}
	ok := canArrange(arr[:], 2)
	fmt.Println(ok)
}
