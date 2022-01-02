package exercise

func productExceptSelf(nums []int) []int {
	l, r := make([]int, len(nums)), make([]int, len(nums))
	l[0] = 1
	r[len(nums)-1] = 1
	for i, k := range nums[:len(nums)-1] {
		l[i+1] = l[i] * k
	}

	for i := len(nums) - 1; i > 0; i-- {
		r[i-1] = r[i] * nums[i]
	}

	for i, k := range r {
		l[i] *= k
	}

	return l
}

// 使用一个int来代替[]int
func productExceptSelf_(nums []int) []int {
	l := make([]int, len(nums))
	l[0] = 1
	r := 1
	for i, k := range nums[:len(nums)-1] {
		l[i+1] = l[i] * k
	}

	for i := len(nums) - 1; i >= 0; i-- {
		l[i] *= r
		r *= nums[i]
	}

	return l
}
