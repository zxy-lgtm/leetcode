package leetcode

func maxProfit(prices []int) int {
	var minV, maxV = prices[0], 0
	for i := 1; i < len(prices); i++ {
		// 当前减去之前的最小值，得到当前位置的最大利益
		// 再比较当前利益是否是最大利益
		if prices[i]-minV > maxV {
			maxV = prices[i] - minV
		}
		// 每次都更新最小值
		if prices[i] < minV {
			minV = prices[i]
		}
	}
	return maxV
}
