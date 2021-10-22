package leetcode

// map的低效做法
func majorityElement(nums []int) []int {

	nums_zip := make(map[int]int)

	for _, value := range nums {
		nums_zip[value]++
	}

	res := make([]int, 0)
	for i, value := range nums_zip {
		if value > len(nums)/3 {
			res = append(res, i)
		}
	}

	return res

}
