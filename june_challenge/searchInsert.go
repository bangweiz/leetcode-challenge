package june_challenge

func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums) - 1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left - 1
}