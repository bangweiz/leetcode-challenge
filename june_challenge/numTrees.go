package june_challenge

func numTrees(n int) int {
	if n <= 2 {
		return n
	}
	res := make([]int, n + 1)
	res[0], res[1] = 1, 1
	for i := 0; i <= n; i++ {
		for j := 1; j <= i; j++ {
			res[0] += res[j - 1] * res[i - j]
		}
	}
	return res[n]
}
