package search

func Search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left := 0
	right := length - 1
	var mid, first int
	for left <= right {
		mid = (right - left) / 2 + left
		first = nums[left]
		if first == target {
			return left
		}

		if nums[mid] < target {
			if first <= nums[mid] {
				left = mid + 1
			} else {
				if first < target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		} else if nums[mid] > target {
			if first > nums[mid] {
				right = mid - 1
			} else {
				if first > target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		} else {
			return mid
		}
	}
	return -1
}