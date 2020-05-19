package numIslands

func numIslands(grid [][]byte) int {
	res := 0
	q := make([]point, 0)
	isVisited := make([][]bool, len(grid))
	for i := range isVisited {
		isVisited[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !isVisited[i][j] && grid[i][j] == '1' {
				res++
				q = append(q, point{i,j})
				wfs(q, isVisited, grid)
			}
		}
	}
	return res
}

func wfs(q []point, isVisited [][]bool, grid [][]byte)  {
	for len(q) != 0 {
		current := q[0]
		q = q[1:]
		if isVisited[current.x][current.y] {
			continue
		}
		isVisited[current.x][current.y] = true

		if current.x + 1 <= len(grid) - 1 {
			if !isVisited[current.x + 1][current.y] && grid[current.x + 1][current.y] == '1' {
				q = append(q, point{current.x + 1, current.y})
			}
		}

		if current.x - 1 >= 0 {
			if !isVisited[current.x - 1][current.y] && grid[current.x - 1][current.y] == '1' {
				q = append(q, point{current.x - 1, current.y})
			}
		}

		if current.y + 1 <= len(grid[current.x]) - 1 {
			if !isVisited[current.x][current.y + 1] && grid[current.x][current.y + 1] == '1' {
				q = append(q, point{current.x, current.y + 1})
			}
		}

		if current.y - 1 >= 0 {
			if !isVisited[current.x][current.y - 1] && grid[current.x][current.y - 1] == '1' {
				q = append(q, point{current.x, current.y - 1})
			}
		}
	}
}

type point struct {
	x int
	y int
}
