package exercise

func numIslands(grid [][]byte) int {
	count := 0
	for i, k := range grid {
		for j, l := range k {
			if l == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, r int, c int) {
	if ok := inArea(grid, r, c); !ok {
		return
	}

	if grid[r][c] != '1' {
		return
	}

	grid[r][c] += 1

	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

func inArea(grid [][]byte, r int, c int) bool {
	if r >= len(grid) || c >= len(grid[0]) || r < 0 || c < 0 {
		return false
	}
	return true
}
