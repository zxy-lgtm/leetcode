package leetcode

func moveZeroes(nums []int) {
	head, tail := 0, 0
	for tail < len(nums)-1 && head < len(nums)-1 {
		if nums[head] != 0 {
			head++
			tail = head
		}
		if nums[tail] == 0 {
			tail++
		}
		if tail == len(nums) {
			break
		}
		if nums[head] == 0 && nums[tail] != 0 {
			nums[tail], nums[head] = nums[head], nums[tail]
		}
	}
}

func moveZeroes2(nums []int) {
	head, tail := 0, 0

	for tail < len(nums) {
		if nums[head] == 0 && nums[tail] != 0 {
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
		}
		tail++
	}
}

/*
func main() {
	nums := []int{0, 1, 0, 9, 7, 2}
	moveZeroes2(nums)
	fmt.Println(nums)
}
*/
