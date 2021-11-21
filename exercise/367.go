package leetcode

func isPerfectSquare(num int) bool {
	l, r := 0, num
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		if mid*mid < num {
			l = mid + 1
		} else if mid*mid > num {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}
