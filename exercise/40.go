package exercise

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var trcak []int
	var res [][]int
	sort.Ints(candidates)
	backtracking(0, 0, target, candidates, trcak, &res)
	return res
}
func backtracking(startIndex, sum, target int, candidates, trcak []int, res *[][]int) {
	//终止条件
	if sum == target {
		tmp := make([]int, len(trcak))
		//拷贝
		copy(tmp, trcak)
		//放入结果集
		*res = append(*res, tmp)
		return
	}
	//回溯
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		// 若当前树层有使用过相同的元素，则跳过
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		//更新路径集合和sum
		trcak = append(trcak, candidates[i])
		sum += candidates[i]
		backtracking(i+1, sum, target, candidates, trcak, res)
		//回溯
		trcak = trcak[:len(trcak)-1]
		sum -= candidates[i]
	}
}
