package new21Game

import "testing"

func TestNew21Game(t *testing.T) {
	res := new21Game(4,3,3)
	t.Logf("Success, the res is %v", res)
}