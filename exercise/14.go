package exercise

func longestCommonPrefix(strs []string) string {
	if len(strs) < 2 {
		return strs[0]
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(res); j++ {
			if len(strs[i]) <= j {
				res = res[:j]
				break
			}
			if res[j] != strs[i][j] {
				res = res[:j]
			}
		}
	}
	return res
}
