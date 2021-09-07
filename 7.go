package leetcode

import "math"

func reverse(x int) int {
	re := 0
	for x != 0 {
		if re < math.MinInt32/10 || re > math.MaxInt32/10 {
			return 0
		}
		re *= 10
		ch := x % 10
		x /= 10
		if ch != 0 {
			re = re + ch
		}

	}
	return re

}
