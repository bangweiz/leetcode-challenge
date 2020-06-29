package maximalSquare

import "testing"

func TestMaximalSquare(t *testing.T) {
	data := make([][]byte, 5)
	data[0] = []byte{'1','0','1','1','1','0'}
	data[1] = []byte{'1','0','1','1','0','0'}
	data[2] = []byte{'1','0','1','1','1','1'}
	data[3] = []byte{'1','0','0','1','1','1'}
	data[4] = []byte{'1','1','1','1','1','1'}

	res := maximalSquare(data)

	if res != 9 {
		t.Errorf("Failed, the res should be %d, but it is %d", 9, res)
	} else {
		t.Log("Success")
	}
}