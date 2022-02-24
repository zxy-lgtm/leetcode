package main

import "fmt"

func next(prt string) (next []int) {
	lp := len(prt)
	next = append(next, -1)
	f := -1
	e := 0
	for e < lp-1 {
		if f == -1 || prt[f] == prt[e] {
			f++
			e++
			//fmt.Println(f)
			next = append(next, f)
		} else {
			f = next[f]
		}
	}

	return
}

func kmp(str, prt string) bool {
	n := next(prt)
	j := -1
	i := 0
	for i < len(str) {
		if j == -1 || prt[j] == str[i] {
			j++
			i++
		} else {
			j = n[j]
		}
	}

	if j == len(prt) {
		return true
	}

	return false
}

func main() {
	prt := "aaab"
	str := "aaaaaaaaab"
	ok := kmp(str, prt)
	fmt.Println(ok)

}
