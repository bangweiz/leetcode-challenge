package june_challenge

func hIndex(citations []int) int {
	if len(citations) == 0 || citations[len(citations) - 1] == 0 {
		return 0
	}
	left, right, mid := 0, len(citations) - 1, 0
	for left < right {
		mid = (left + right) / 2
		if len(citations) - mid < citations[mid] {
			right = mid - 1
		} else if len(citations) - mid > citations[mid] {
			left = mid + 1
		} else {
			return citations[mid]
		}
	}
	value := citations[left]
	if len(citations) - left > value {
		return len(citations) - left - 1
	} else if len(citations) - left < value {
		return len(citations) - left
	}
	return value
}
