package exercise

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	tag := 0
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
		dp[i][0] = 1
		if obstacleGrid[i][0] == 1 {
			tag = 1
		}
		if tag == 1 {
			dp[i][0] = 0
		}
	}

	tag = 0
	for i := 0; i < len(obstacleGrid[0]); i++ {
		dp[0][i] = 1
		if obstacleGrid[0][i] == 1 {
			tag = 1
		}
		if tag == 1 {
			dp[0][i] = 0
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			if obstacleGrid[i][j] != 0 {
				dp[i][j] = 0
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

// 滚动数组优化
func uniquePathsWithObstacles_(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}
