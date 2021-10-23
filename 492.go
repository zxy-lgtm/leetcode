package leetcode

import "math"

func constructRectangle(area int) []int {

	res := int(math.Sqrt(float64(area)))

	for area%res != 0 {
		res--
	}

	return []int{area / res, res}

}
