package exercise

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
