package exercise

func shortestPathBinaryMatrix(grid [][]int) int {
	rows := len(grid)
	if grid == nil || rows == 0 || grid[0][0] == 1 || grid[rows-1][rows-1] == 1 {
		return -1
	}
	if len(grid) == 1 {
		return 1
	}
	direction := []int{-1, 0, 1}
	grid[0][0] = 1
	//途经的每一个点都记录从起点到次的长度
	que := make([]int, 0)
	que = append(que, 0)
	//用que记录当前点的坐标，判断有没有下一个节点

	var x, y, nx, ny int

	for len(que) > 0 {
		x, y = que[0]/rows, que[0]%rows
		que = que[1:]
		for _, i := range direction {
			for _, j := range direction {
				if i == j && i == 0 {
					continue
				}
				nx, ny = x+i, y+j
				if nx < rows && ny < rows && nx >= 0 && ny >= 0 && grid[nx][ny] == 0 {
					que = append(que, nx*rows+ny)
					grid[nx][ny] = grid[x][y] + 1
					if nx == rows-1 && ny == rows-1 {
						return grid[nx][ny]
					}
				}
			}
		}
	}
	return -1
}
