package findMedianSortedArrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1,9,11,12,17,18}
	nums2 := []int{5,8,9,15,16,19}

	res := findMedianSortedArrays(nums1, nums2)

	if res != 11.5 {
		t.Errorf("Failed, the res should be %v, bu it is %v", 11.5, res)
	} else {
		t.Logf("Success")
	}
}