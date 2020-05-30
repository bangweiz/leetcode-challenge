package findMedianSortedArrays

// Time complexity O(m + n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num := make([]int, 0)
	j, k := 0, 0
	for j < len(nums1) && k < len(nums2) {
		if nums1[j] < nums2[k] {
			num = append(num, nums1[j])
			j++
		} else {
			num = append(num, nums2[k])
			k++
		}
	}

	for j < len(nums1) {
		num = append(num, nums1[j])
		j++
	}

	for k < len(nums2) {
		num = append(num, nums2[k])
		k++
	}

	if len(num) % 2 == 0 {
		return float64(num[len(num) / 2] + num[len(num) / 2 - 1]) / 2
	} else {
		return float64(num[len(num) / 2])
	}
}