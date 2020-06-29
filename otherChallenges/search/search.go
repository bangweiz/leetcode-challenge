package search

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left := 0
	right := length - 1
	var mid int
	for left <= right {
		mid = (right - left) / 2 + left

		if nums[mid] < target {
			if nums[left] <= nums[mid] {
				left = mid + 1
			} else {
				if nums[left] < target {
					right = mid - 1
				} else if nums[left] > target {
					left = mid + 1
				} else {
					return left
				}
			}
		} else if nums[mid] > target {
			if nums[left] > nums[mid] {
				right = mid - 1
			} else {
				if nums[left] > target {
					left = mid + 1
				} else if nums[left] < target {
					right = mid - 1
				} else {
					return left
				}
			}
		} else {
			return mid
		}
	}
	return -1
}