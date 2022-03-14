package exercise

func climbStairs(n int) int {
	dp_1 := 1
	dp_2 := 2

	if n <= 3 {
		return n
	}

	for i := 3; i <= n; i++ {
		tmp := dp_1 + dp_2
		dp_1 = dp_2
		dp_2 = tmp
	}

	return dp_2
}
