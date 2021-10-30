package leetcode

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	min := xorSum & (-xorSum)
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&min > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}
