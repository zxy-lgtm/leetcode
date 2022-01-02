package exercise

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	small, mid := math.MaxInt32, math.MaxInt32
	for _, k := range nums {
		if k <= small {
			small = k
		} else if k <= mid {
			mid = k
		} else {
			return true
		}
	}

	return false
}
