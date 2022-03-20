package exercise

func maxProfit(prices []int) int {
	dp0, dp1, dp2 := -prices[0], 0, 0

	for _, v := range prices {
		dp0, dp1, dp2 = max(dp0, dp2-v), dp0+v, max(dp1, dp2)
	}

	return max(dp1, dp2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
