package solveSudoku

import "strconv"

func solveSudoku(board [][]byte)  {
	points := make([]*point, 0)
	row, col, squ := make([]map[byte]int, 9), make([]map[byte]int, 9), make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		row[i], col[i], squ[i] = make(map[byte]int), make(map[byte]int), make(map[byte]int)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				row[i][board[i][j]] = 1
				col[j][board[i][j]] = 1
				squ[(i / 3) * 3 + j / 3][board[i][j]] = 1
			} else {
				points = append(points, &point{i, j})
			}
		}
	}

	fill(points, row, col, squ, 0, board)
}

func fill(points []*point, row, col, squ []map[byte]int, n int, board [][]byte) bool {
	current := points[n]
	for i := 1; i < 10; i++ {
		b := []byte(strconv.Itoa(i))[0]
		_, rowOk := row[current.i][b]
		_, colOk := col[current.j][b]
		_, squOk := squ[(current.i / 3) * 3 + current.j / 3][b]
		if !rowOk && !colOk && !squOk {
			row[current.i][b] = 1
			col[current.j][b] = 1
			squ[(current.i / 3) * 3 + current.j / 3][b] = 1
			board[current.i][current.j] = b
			if n + 1 == len(points) {
				return true
			} else {
				res := fill(points, row, col, squ, n + 1, board)
				if res {
					return true
				}
				delete(row[current.i], b)
				delete(col[current.j], b)
				delete(squ[(current.i / 3) * 3 + current.j / 3], b)
			}
		}
	}
	return false
}

type point struct {
	i int
	j int
}