package largestNumber

import "testing"

func TestLargestNumber(t *testing.T) {
	cost := []int{4,3,2,5,6,7,2,5,5}

	res := largestNumber(cost, 9)

	if res != "7772" {
		t.Errorf("Failed, the result should be %v, but it is %v", "7772", res)
	} else {
		t.Log("Success")
	}
}