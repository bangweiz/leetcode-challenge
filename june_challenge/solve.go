package june_challenge

type point struct {
	x int
	y int
}

func solve(board [][]byte)  {
	if len(board) == 0 {
		return
	}
	isVisited := make([][]bool, len(board))
	data := make([][]byte, len(board))
	for i := range board {
		isVisited[i] = make([]bool, len(board[i]))
		data[i] = make([]byte, len(board[i]))
		for j := range data[i] {
			data[i][j] = 'X'
		}
	}
	q := make([]point, 0)
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			q = append(q, point{0, i})
		}
		if board[len(board) - 1][i] == 'O' {
			q = append(q, point{len(board) - 1, i})
		}
	}
	for i := 1; i < len(board) - 1; i++ {
		if board[i][0] == 'O' {
			q = append(q, point{i, 0})
		}
		if board[i][len(board[i]) - 1] == 'O' {
			q = append(q, point{i, len(board[i]) - 1})
		}
	}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		if !isVisited[current.x][current.y] {
			isVisited[current.x][current.y] = true
			data[current.x][current.y] = 'O'
			if current.x + 1 <= len(board) - 1 {
				if !isVisited[current.x + 1][current.y] && board[current.x + 1][current.y] == 'O' {
					q = append(q, point{current.x + 1, current.y})
				}
			}

			if current.x - 1 >= 0 {
				if !isVisited[current.x - 1][current.y] && board[current.x - 1][current.y] == 'O' {
					q = append(q, point{current.x - 1, current.y})
				}
			}

			if current.y + 1 <= len(board[current.x]) - 1 {
				if !isVisited[current.x][current.y + 1] && board[current.x][current.y + 1] == 'O' {
					q = append(q, point{current.x, current.y + 1})
				}
			}

			if current.y - 1 >= 0 {
				if !isVisited[current.x][current.y - 1] && board[current.x][current.y - 1] == 'O' {
					q = append(q, point{current.x, current.y - 1})
				}
			}
		}
	}

	copy(board, data)
}