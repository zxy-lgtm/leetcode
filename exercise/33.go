package exercise

import "sort"

func search(nums []int, target int) int {
	res := sort.Search(len(nums), func(i int) bool {
		if nums[i] >= nums[0] {
			return target >= nums[0] && nums[i] >= target
		} else {
			return target >= nums[0] || nums[i] >= target
		}
	})

	if res < len(nums) && nums[res] == target {
		return res
	}
	return -1
}

/*func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}*/
