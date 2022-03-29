package exercise

// 暴力法
func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide
	}
	rows, columns := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				maxSide = max(maxSide, 1)
				curMaxSide := min(rows-i, columns-j)
				for k := 1; k < curMaxSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for m := 0; m < k; m++ {
						if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 动态规划
func maximalSquare_(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	res := 0
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		dp[i][0] = int(matrix[i][0] - '0')
		res = max(dp[i][0], res)
	}

	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = int(matrix[0][i] - '0')
		res = max(dp[0][i], res)
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				res = max(dp[i][j], res)
			}
		}
	}
	return res * res
}
