package minPathSum

func MinPathSum(grid [][]int) int {
	res := make([][]int, len(grid))
	for i := range res {
		res[i] = make([]int, len(grid[0]))
	}
	res[0][0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		res[0][i] = grid[0][i] + res[0][i - 1]
	}

	for i := 1; i < len(grid); i++ {
		res[i][0] = res[i - 1][0] + grid[i][0]
		for j := 1; j < len(grid[i]); j++ {
			if res[i - 1][j] > res[i][j - 1] {
				res[i][j] = res[i][j - 1] + grid[i][j]
			} else {
				res[i][j] = res[i - 1][j] + grid[i][j]
			}
		}
	}

	return res[len(res) - 1][len(res[0]) - 1]
}