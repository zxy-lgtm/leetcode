package exercise

func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	temp := 1
	if k <= 1 {
		return 0
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		temp *= nums[j]
		for i <= j && temp >= k {
			temp /= nums[i]
			i++
		}
		ans += j - i + 1
	}
	return ans
}
