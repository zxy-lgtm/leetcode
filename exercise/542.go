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

func updateMatrix2(matrix [][]int) [][]int {

	n, m := len(matrix), len(matrix[0])
	queue := make([][]int, 0)
	for i := 0; i < n; i++ { // 把0全部存进队列，后面从队列中取出来，判断每个访问过的节点的上下左右，直到所有的节点都被访问过为止。
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				point := []int{i, j}
				queue = append(queue, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 { // 这里就是 BFS 模板操作了。
		point := queue[0]
		queue = queue[1:]
		for _, v := range direction {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}

	return matrix
}
