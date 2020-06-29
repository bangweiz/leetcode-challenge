package june_challenge

func largestDivisibleSubset(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	nums = MergeSort(nums)
	dp, res, max := make([]int, len(nums)), make([]int, len(nums)), 0

	for i := 0; i < len(dp); i++ {
		for j := i; j >= 0; j-- {
			if nums[i] % nums[j] == 0 {
				if dp[j] + 1 > dp[i] {
					dp[i], res[i] = dp[j] + 1, j
				}
			}
		}
		if dp[i] > dp[max] {
			max = i
		}
	}
	r := make([]int, 0)
	for res[max] != max {
		r = append(r, nums[max])
		max = res[max]
	}
	r = append(r, nums[max])
	return r
}

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return Merge(MergeSort(nums[:mid]), MergeSort(nums[mid:]))
}

func Merge(a, b []int) []int {
	c, i, j, index := make([]int, len(a) + len(b)), 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[index] = a[i]
			i++
		} else {
			c[index] = b[j]
			j++
		}
		index++
	}

	for i < len(a) {
		c[index] = a[i]
		i++
		index++
	}

	for j < len(b) {
		c[index] = b[j]
		j++
		index++
	}
	return c
}
