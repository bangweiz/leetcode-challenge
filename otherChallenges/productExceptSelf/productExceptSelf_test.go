package productExceptSelf

import "testing"

func TestProductExceptSelf(t *testing.T) {
	data := []int{1,2,3,4}
	res := productExceptSelf(data)

	if len(res) != len(data) || res[0] != 24 || res[1] != 12 || res[2] != 8 || res[3] != 6 {
		t.Error("Failed")
	} else {
		t.Log("Success")
	}
}