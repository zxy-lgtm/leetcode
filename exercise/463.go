package exercise

func islandPerimeter(grid [][]int) (res int) {
	for i, nums := range grid {
		for j, k := range nums {
			if k == 1 {
				res += getl(grid, i, j)
			}
		}
	}
	return
}

func getl(grid [][]int, r int, l int) int {
	res := 4
	if r-1 >= 0 && grid[r-1][l] == 1 {
		res--
	}
	if l-1 >= 0 && grid[r][l-1] == 1 {
		res--
	}
	if r+1 < len(grid) && grid[r+1][l] == 1 {
		res--
	}
	if l+1 < len(grid[0]) && grid[r][l+1] == 1 {
		res--
	}
	return res
}
