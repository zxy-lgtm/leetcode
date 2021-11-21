package leetcode

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	re := make([]string, numRows)
	n := 2*numRows - 2
	for i, char := range s {
		x := i % n
		if (n - x) > x {
			re[x] += string(char)
		} else {
			re[n-x] += string(char)
		}
	}
	return strings.Join(re, "")
}
