package exercise

import "math"

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange(coins []int, amount int) int {
	mins := make([]int, amount+1)
	var dp func(coins []int, amount int) int
	dp = func(coins []int, amount int) int {
		if amount == 0 {
			return 0
		}

		if amount < 0 {
			return -1
		}

		if mins[amount] != 0 {
			return mins[amount]
		}

		res := int(^uint(0) >> 1)
		for _, coin := range coins {
			min := dp(coins, amount-coin)
			if min != -1 {
				res = Min(res, min+1)
			}
		}
		if res == int(^uint(0)>>1) {
			mins[amount] = -1
		} else {
			mins[amount] = res
		}

		return mins[amount]
	}

	return dp(coins, amount)
}

func coinChange_(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 初始化为math.MaxInt32
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				//fmt.Println(dp,j,i)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
