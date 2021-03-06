package leetcode

func mySqrt(x int) int {
	left, right := 0, x
	mid := 0

	for left <= right {
		mid = left + (right-left)/2
		if mid*mid <= x {
			if (mid+1)*(mid+1) > x {
				return mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return mid
}
