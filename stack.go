package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	stack := []int{}
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack)-1 >= 0 && stack[len(stack)-1] == popped[i] {
			i++
			stack = stack[:len(stack)-1]
		}
		if i == n {
			break
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

//2231 最小栈
type MinStack struct {
	stack1, stack2 []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack1: []int{},
		//stack2维护一个元素递减的栈
		stack2: []int{},
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack1) == 0 || this.stack2[len(this.stack2)-1] > x {
		this.stack2 = append(this.stack2, x)
	}
	this.stack1 = append(this.stack1, x)
}

func (this *MinStack) Pop() {
	if this.stack1[len(this.stack1)-1] == this.stack2[len(this.stack2)-1] {
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	this.stack1 = this.stack1[:len(this.stack1)-1]
}

func (this *MinStack) Top() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) Min() int {

	return this.stack2[len(this.stack2)-1]
}

