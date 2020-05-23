package moveZeroes

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{1,2,3,0,0,4,0,5,0}
	res := []int{1,2,3,4,5,0,0,0,0}
	moveZeroes(nums)

	for i := range res {
		if res[i] != nums[i] {
			t.Errorf("Failed")
			return
		}
	}
	t.Log("Success")
}
