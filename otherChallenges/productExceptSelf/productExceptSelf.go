package productExceptSelf

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[len(nums) - 1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		res[i] = res[i + 1] * nums[i + 1]
	}
	left := 1
	for i := range nums {
		res[i], left = left * res[i], left * nums[i]
	}
	return res
}