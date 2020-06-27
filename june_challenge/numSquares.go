package june_challenge

func numSquares(n int) int {
	nums := []int{0}
	for i := 1; true; i++ {
		temp := i * i
		if temp > n {
			break
		}
		nums = append(nums, temp)
	}
	
	res := make([]int, n + 1)
	res[1] = 1
	for i := 2; i <= n; i++ {
		temp := 999999999
		for j := 1; j < len(nums); j++ {
			if i - nums[j] < 0 {
				break
			} else {
				temp = min(temp, res[i - nums[j]] + 1)
			}
		}
		res[i] = temp
	}

	return res[n]
}
