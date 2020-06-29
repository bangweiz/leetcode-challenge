package solveSudoku

import (
	"fmt"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	board := make([][]byte, 9)
	board[0] = []byte{'.','8','5','.','.','4','.','.','.'}
	board[1] = []byte{'7','.','.','.','.','.','.','9','.'}
	board[2] = []byte{'.','1','.','.','.','.','7','6','.'}
	board[3] = []byte{'.','.','.','.','.','.','.','.','.'}
	board[4] = []byte{'.','9','.','4','5','2','.','8','3'}
	board[5] = []byte{'.','4','.','.','8','.','.','.','5'}
	board[6] = []byte{'.','.','6','.','.','.','.','7','.'}
	board[7] = []byte{'.','.','.','6','.','.','.','.','.'}
	board[8] = []byte{'4','.','3','.','.','8','.','.','.'}

	solveSudoku(board)

	for _, v := range board {
		fmt.Printf("%s ", v)
		fmt.Println()
	}
}