package exercise

func maxProfit(prices []int) int {
	dp0, dp1 := 0, -prices[0]

	for _, v := range prices {
		dp0, dp1 = max(dp0, dp1+v), max(dp1, dp0-v)
	}

	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
