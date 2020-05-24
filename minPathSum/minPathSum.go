package minPathSum

func minPathSum(grid [][]int) int {
	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] = grid[0][i] + grid[0][i - 1]
	}

	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i - 1][0] + grid[i][0]
		for j := 1; j < len(grid[i]); j++ {
			if grid[i - 1][j] > grid[i][j - 1] {
				grid[i][j] = grid[i][j - 1] + grid[i][j]
			} else {
				grid[i][j] = grid[i - 1][j] + grid[i][j]
			}
		}
	}

	return grid[len(grid) - 1][len(grid[0]) - 1]
}