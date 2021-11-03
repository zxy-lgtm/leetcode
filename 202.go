package leetcode

func isHappy(n int) bool {
	all_nums := make(map[int]int)
	r := Sum(n)
	for r != 1 {
		if all_nums[r] != 0 {
			return false
		}
		all_nums[r]++
		n = r
		r = Sum(n)
	}
	return true

}

func Sum(num int) (sum int) {
	for num != 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	return
}
