package minStack

type MinStack struct {
	stack []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 || x <= this.min[len(this.min) - 1] {
		this.min = append(this.min, x)
	}
}


func (this *MinStack) Pop()  {
	if this.min[len(this.min) - 1] == this.stack[len(this.stack) - 1] {
		this.min = this.min[: len(this.min) - 1]
	}
	this.stack = this.stack[: len(this.stack) - 1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.min[len(this.min) - 1]
}