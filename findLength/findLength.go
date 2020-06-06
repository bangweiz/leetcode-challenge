package findLength

func findLength(A []int, B []int) int {
	res, m := make([][]int, len(A) + 1), 0
	for i := range res {
		res[i] = make([]int, len(B) + 1)
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				res[i + 1][j + 1] = res[i][j] + 1
				m = max(res[i + 1][j + 1], m)
			}
		}
	}
	return m
}

func max(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
