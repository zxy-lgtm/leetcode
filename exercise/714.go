package exercise

func maxProfit(prices []int, fee int) int {
	dp0, dp1 := -prices[0], 0

	for _, v := range prices {
		dp0, dp1 = max(dp0, dp1-v), max(dp1, dp0+v-fee)
	}
	return dp1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
