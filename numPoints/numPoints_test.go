package numPoints

import "testing"

func TestNumPoints(t *testing.T) {
	points := make([][]int, 5)
	points[0] = []int{5, 5}
	points[1] = []int{-2, 5}
	points[2] = []int{4, 2}
	points[3] = []int{1, -1}
	points[4] = []int{1, 1}

	res := numPoints(points, 5)

	if res != 5 {
		t.Errorf("Failed, the result should be %d, but it is %d", 5, res)
	} else {
		t.Log("Success")
	}
}