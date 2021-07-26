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

/*
func main() {
	nums := []int{1, 2, 3}
	rotate(nums, 4)
	fmt.Println(nums)
}
*/
