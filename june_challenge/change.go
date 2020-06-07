package june_challenge

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	res := make([][]int, len(coins) + 1)
	for i := range res {
		res[i] = make([]int, amount + 1)
	}
	for i := 1; i < len(res); i++ {
		for j := 1; j < len(res[i]); j++ {
			if j < coins[i - 1] {
				res[i][j] = res[i - 1][j]
			} else if j > coins[i - 1] {
				res[i][j] = res[i][j - coins[i - 1]] + res[i - 1][j]
			} else {
				res[i][j] = res[i - 1][j] + 1
			}
		}
	}

	return res[len(res) - 1][amount]
}