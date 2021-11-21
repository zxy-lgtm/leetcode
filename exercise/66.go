package leetcode

func plusOne(digits []int) []int {
	len := len(digits)

	if digits[len-1] != 9 {
		digits[len-1]++
		return digits
	}

	i := len - 1

	for digits[i] == 9 {
		digits[i] = 0
		i--
		if i < 0 {
			digits = append([]int{1}, digits...)
			return digits
		}
	}

	digits[i]++
	return digits
}
