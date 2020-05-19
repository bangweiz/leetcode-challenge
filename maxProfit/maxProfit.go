package maxProfit

func MaxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if n := prices[i] - prices[i - 1]; n > 0 {
			res += n
		}
	}
	return res
}