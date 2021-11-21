package leetcode

func twoSum(numbers []int, target int) []int {
	head, tail := 0, len(numbers)-1
	for head < tail {
		if (numbers[head] + numbers[tail]) == target {
			return []int{head + 1, tail + 1}
		}
		if (numbers[head] + numbers[tail]) > target {
			tail--
		}

		if (numbers[head] + numbers[tail]) < target {
			head++
		}
	}
	return []int{0, 0}
}
