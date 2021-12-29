package exercise

func singleNumber(nums []int) (res int) {
	for _, k := range nums {
		res ^= k
	}

	return res
}
