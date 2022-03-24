package exercise

func generate(numRows int) [][]int {
	ans := make([][]int, 0)
	for i := 1; i <= numRows; i++ {
		tmp := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 {
				tmp[j] = 1
			} else if j >= i-1 {
				tmp[j] = 1
			} else {
				tmp[j] = ans[i-2][j-1] + ans[i-2][j]
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}
