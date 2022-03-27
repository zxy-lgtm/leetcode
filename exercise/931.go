package exercise

func minFallingPathSum(matrix [][]int) int {
	//初始化memo
	memo := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		memo[i] = make([]int, len(matrix))
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = 100000
		}
	}

	for j := 0; j < len(matrix); j++ {
		memo[0][j] = matrix[0][j]
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		//判断下表是否合法
		if i < 0 || i >= len(matrix[0]) || j < 0 || j >= len(matrix[0]) {
			return 100000
		}
		//base case
		if i == 0 {
			return matrix[0][j]
		}
		//判断是否已在备忘录中存在
		if memo[i][j] != 100000 {
			return memo[i][j]
		}
		//不在备忘录中存在，则计算
		memo[i][j] = matrix[i][j] + min(dp(i-1, j), dp(i-1, j-1), dp(i-1, j+1))

		return memo[i][j]
	}

	//寻找最小值
	res := 100000
	for i := 0; i < len(matrix); i++ {
		res = min(res, dp(len(matrix)-1, i))
	}
	return res
}

func min(nums ...int) int {
	min := 100000
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}
