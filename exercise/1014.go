package exercise

func maxScoreSightseeingPair(values []int) int {
	ans := 0
	maxnum := values[0]
	for i := 1; i < len(values); i++ {
		ans = max(ans, maxnum+values[i]-i)
		maxnum = max(maxnum, values[i]+i)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
