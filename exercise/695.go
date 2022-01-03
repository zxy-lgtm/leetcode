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

func maxAreaOfIsland2(grid [][]int) (result int) {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var bfs func(i, j int) int
	bfs = func(i, j int) (count int) {
		arr := [][2]int{{i, j}}
		for len(arr) != 0 {
			a, b := arr[0][0], arr[0][1]
			arr = arr[1:]
			if a >= 0 && a < n && b >= 0 && b < m && grid[a][b] == 1 {
				count++
				grid[a][b] = 0
				arr = append(arr, [][2]int{{a - 1, b}, {a, b - 1}, {a + 1, b}, {a, b + 1}}...)
			}
		}
		return
	}
	for i := range grid {
		for j := range grid[0] {
			result = Max(result, bfs(i, j))
		}
	}
	return
}

//Max return the maximum number
func Max(a ...int) int {
	max := a[0]
	for _, v := range a[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// rewrite
func maxAreaOfIsland_(grid [][]int) int {
	max := 0
	for i, nums := range grid {
		for j, k := range nums {
			if k == 1 {
				area := dfs_(grid, i, j)
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}

func dfs_(grid [][]int, r int, c int) int {
	if ok := inarea(grid, r, c); !ok {
		return 0
	}
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c]++

	return 1 + dfs_(grid, r-1, c) + dfs_(grid, r+1, c) + dfs_(grid, r, c-1) + dfs_(grid, r, c+1)
}

func inarea(grid [][]int, r int, c int) bool {
	if r >= len(grid) || c >= len(grid[0]) || r < 0 || c < 0 {
		return false
	}
	return true
}
