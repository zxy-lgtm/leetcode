package leetcode

func rotate(nums []int, k int) {
	l := len(nums)
	ans := make([]int, l)
	if k > l {
		k = k % l
	}
	index := l - k

	for i := 0; i < k; i++ {
		ans[i] = nums[index+i]
		//fmt.Println(ans[i])
	}

	for i := 0; i < index; i++ {
		ans[i+k] = nums[i]
	}

	for i := 0; i < l; i++ {
		nums[i] = ans[i]
	}

}

//环状替换
func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

//翻转数组
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

/*
func main() {
	nums := []int{1, 2, 3}
	rotate(nums, 4)
	fmt.Println(nums)
}
*/
