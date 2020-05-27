package findMaxLength

import "testing"

func TestFindMaxLength(t *testing.T)  {
	data := []int{0,1}
	res := findMaxLength(data)
	if res != 2 {
		t.Errorf("Failed, the res should be %v, but it is %v", 2, res)
	} else {
		t.Log("Success")
	}
}