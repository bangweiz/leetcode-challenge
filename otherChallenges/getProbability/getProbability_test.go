package getProbability

import (
	"fmt"
	"testing"
)

func TestGetProbability(t *testing.T) {
	data := []int{6,6,6,6,6,6}
	res := getProbability(data)
	fmt.Println(res)
	t.Log("Success")
}
