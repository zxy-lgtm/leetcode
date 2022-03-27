package exercise

func matrixBlockSum(mat [][]int, k int) [][]int {
	m := len(mat)
	n := len(mat[0])
	answer := make([][]int, m)
	for i := 0; i < m; i++ {
		answer[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			answer[i][j] = calculate(i, j, mat, k)
		}
	}

	return answer
}
func calculate(i, j int, mat [][]int, k int) int {
	rMin := i - k
	if rMin < 0 {
		rMin = 0
	}
	rMax := i + k
	if rMax > len(mat)-1 {
		rMax = len(mat) - 1
	}

	cMin := j - k
	if cMin < 0 {
		cMin = 0
	}
	cMax := j + k
	if cMax > len(mat[0])-1 {
		cMax = len(mat[0]) - 1
	}
	sum := 0
	for m := rMin; m <= rMax; m++ {
		for n := cMin; n <= cMax; n++ {
			sum += mat[m][n]
		}
	}
	return sum
}
