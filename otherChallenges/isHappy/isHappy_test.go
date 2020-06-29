package isHappy

import "testing"

func TestIsHappy(t *testing.T) {
	res := isHappy(113)
	if !res {
		t.Log("Success")
	} else {
		t.Errorf("Failed, this res should be %v, but it is %v", !res, res)
	}
}