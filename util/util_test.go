package util

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{6,5,2,9,14,54,23,14,78,32,11,0,65,85,41,23}
	nums = mergeSort(nums)
	for _, v := range nums {
		fmt.Printf("%v ", v)
	}
	t.Log("Success")
}