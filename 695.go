package leetcode

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j, grid))
			}
		}
	}
	return ans
}

func dfs(x, y int, grid [][]int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}
	ans := 1
	grid[x][y] = 0
	ans += dfs(x-1, y, grid)
	ans += dfs(x, y-1, grid)
	ans += dfs(x+1, y, grid)
	ans += dfs(x, y+1, grid)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
