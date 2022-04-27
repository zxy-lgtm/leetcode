package exercise

func partition(s string) [][]string {
	var tmpString []string // 切割字符串集合
	var res [][]string     // 结果集合
	backTracking(s, tmpString, 0, &res)
	return res
}
func backTracking(s string, tmpString []string, startIndex int, res *[][]string) {
	if startIndex == len(s) { //到达字符串末尾了
		t := make([]string, len(tmpString))
		copy(t, tmpString)
		*res = append(*res, t)
	}
	for i := startIndex; i < len(s); i++ {
		if isPartition(s, startIndex, i) {
			tmpString = append(tmpString, s[startIndex:i+1])
		} else {
			continue
		}
		// 递归
		backTracking(s, tmpString, i+1, res)
		// 回溯
		tmpString = tmpString[:len(tmpString)-1]
	}
}

func isPartition(s string, startIndex, end int) bool {
	left := startIndex
	right := end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
