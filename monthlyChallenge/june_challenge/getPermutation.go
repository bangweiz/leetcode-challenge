package june_challenge

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	res := ""
	nums, num := []int{1,1,2,6,24,120,720,5040,40320,362880}, make([]int, n)
	for i := 1; i <= n; i++ {
		num[i - 1] = i
	}
	for i := n - 1; i >= 0; i-- {
		index := k / nums[i]
		k = k % nums[i]
		if k == 0 {
			res += strconv.Itoa(num[index - 1])
			num = append(num[: index - 1], num[index:]...)
			for i := len(num) - 1; i >= 0; i-- {
				res += strconv.Itoa(num[i])
			}
			return res
		}
		res += strconv.Itoa(num[index])
		num = append(num[: index], num[index + 1:]...)

	}
	return res
}
