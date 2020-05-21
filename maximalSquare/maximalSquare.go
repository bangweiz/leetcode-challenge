package maximalSquare

import "math"

func maximalSquare(matrix [][]byte) int {
	max := 0
	if len(matrix) == 0 {
		return max
	}
	data := make([][]int, len(matrix) + 1)
	for i := range data {
		data[i] = make([]int, len(matrix[0]) + 1)
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				data[i + 1][j + 1] = 0
			} else {
				if data[i][j] == 0 || data[i + 1][j] == 0 || data[i][j + 1] == 0 {
					data[i + 1][j + 1] = 1
					if max == 0 {
						max = 1
					}
				} else {
					temp := data[i][j]
					if data[i + 1][j] < temp {
						temp = data[i + 1][j]
					}
					if data[i][j + 1] < temp {
						temp = data[i][j + 1]
					}
					temp = int(math.Sqrt(float64(temp)))
					data[i + 1][j + 1] = (1 + temp) * (1 + temp)
					if data[i + 1][j + 1] > max {
						max = data[i + 1][j + 1]
					}
				}
			}
		}
	}
	return max
}