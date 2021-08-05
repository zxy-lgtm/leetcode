package leetcode

func updateMatrix(mat [][]int) [][]int {
	for i, m := range mat {
		for j, k := range m {
			if k != 0 {
				findzero(mat, i, j)
				mat[i][j]--
			}
		}
	}
	return mat

}

func findzero(mat [][]int, i, j int) {
	if mat[i][j] == 0 {
		return
	}

	mat[i][j]++

	if i > 0 && j > 0 && i < len(mat)-1 && j < len(mat[0])-1 {
		findzero(mat, i+1, j)
		findzero(mat, i-1, j)
		findzero(mat, i, j+1)
		findzero(mat, i, j-1)
	}
}
