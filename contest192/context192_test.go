package contest192

import "testing"

func TestMinCost(t *testing.T) {
	house := []int{3,1,2,3}
	costs := make([][]int, 4)
	costs[0] = []int{1,1,1}
	costs[1] = []int{1,1,1}
	costs[2] = []int{1,1,1}
	costs[3] = []int{1,1,1}

	res := minCost(house, costs, 4, 3, 3)

	if res != -1 {
		t.Errorf("Failed, the result should be %v, but it is %v.", 9, res)
	} else {
		t.Logf("Success, the result is %v.", res)
	}
}