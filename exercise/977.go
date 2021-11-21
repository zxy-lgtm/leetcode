package leetcode

import (
	"math"
	"sort"
)

func sortedSquares1(nums []int) []int {
	tag := 0
	max := int(math.Abs(float64(nums[0])))
	for i := 0; nums[i] < max && i < len(nums)-1; i++ {
		if nums[i] < 0 {
			nums[i] = int(math.Abs(float64(nums[i])))
		}
		tag++
	}
	if tag < len(nums) {
		tag++
	}
	sort.Ints(nums[:tag])
	//fmt.Println(tag, nums)

	for i, _ := range nums {
		nums[i] = int(math.Pow(float64(nums[i]), 2))
	}

	return nums
}

func sortedSquares2(nums []int) []int {
	l := len(nums)
	lastindex := -1
	for i := 0; nums[i] < 0 && i < l; i++ {
		lastindex++
	}
	ans := make([]int, 0, l)
	for i, j := lastindex, lastindex+1; i >= 0 || j < l; {
		if i < 0 {
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == l {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] > nums[j]*nums[j] {
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if nums[i]*nums[i] <= nums[j]*nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i--
		}
	}
	return ans
}

func sortedSquares3(nums []int) []int {
	tail := len(nums) - 1
	head := 0
	if head == tail {
		nums[0] = nums[0] * nums[0]
		return nums
	}

	ans := make([]int, tail+1)
	pos := tail

	for pos >= 0 {
		if nums[head]*nums[head] > nums[tail]*nums[tail] {
			ans[pos] = nums[head] * nums[head]
			pos--
			head++
		} else {
			ans[pos] = nums[tail] * nums[tail]
			pos--
			tail--
		}
	}

	return ans

}

/*func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}*/
