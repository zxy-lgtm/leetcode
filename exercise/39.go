package exercise

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	backtrack(candidates, []int{}, 0, target)
	return res
}

func backtrack(candidates, temp []int, start, target int) {
	if sum(temp) == target {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res = append(res, tmp)
	}
	if sum(temp) > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		temp = append(temp, candidates[i])
		backtrack(candidates, temp, i, target)
		temp = temp[:len(temp)-1]
	}
}
