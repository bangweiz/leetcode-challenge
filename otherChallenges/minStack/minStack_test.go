package minStack

import "testing"

func TestMinStack(t *testing.T) {
	res := make([]int, 0)
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	res = append(res, minStack.GetMin())
	minStack.Pop()
	res = append(res, minStack.Top())
	res = append(res, minStack.GetMin())

	data := []int{-3,0,-2}
	flag := true

	for i := range res {
		if res[i] != data[i] {
			flag = false
			return
		}
	}

	if flag {
		t.Log("Success")
	} else {
		t.Error("Failed")
	}
}