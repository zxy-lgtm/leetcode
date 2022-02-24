package exercise

func fib(n int) int {
	memo := make([]int, n+1)
	return note(memo, n)
}

func note(memo []int, n int) int {
	if n == 0 || n == 1 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = note(memo, n-1) + note(memo, n-2)

	return memo[n]
}

func fib_(n int) int {
	dp_1 := 0
	dp_2 := 1
	if n < 2 {
		return n
	}
	for i := 2; i < n+1; i++ {
		tmp := dp_1 + dp_2
		dp_1 = dp_2
		dp_2 = tmp
	}

	return dp_2
}
