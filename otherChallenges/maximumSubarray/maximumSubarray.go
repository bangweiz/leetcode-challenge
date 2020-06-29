package maximumSubarray

func MaximumSubarray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if max <= 0 {
			max = nums[i]
		} else {
			max += nums[i]
		}
		res = getMax(max, res)
	}
	return getMax(max, res)
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
