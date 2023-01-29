package main

import "fmt"

type MinStack struct {
	Elements []int
	MinS []int
}

func Constructor() MinStack {
	return MinStack{}	
}

func (stack *MinStack) Push(x int) {
	stack.Elements = append(stack.Elements, x)
	if len(stack.MinS) == 0 {
		stack.MinS = append(stack.MinS, x)
	} else {
		min := stack.Min()
		if x < min {
			stack.MinS = append(stack.MinS, x)
		} else {
			stack.MinS = append(stack.MinS, min)
		}
	}
}

func (stack *MinStack) Pop() {
	if len(stack.Elements) == 0 {
		panic("empty stack")
	}
	stack.Elements = stack.Elements[:len(stack.Elements)-1]
	stack.MinS = stack.MinS[:len(stack.MinS)-1]
}

func (stack *MinStack) Top() int {
	n := len(stack.Elements)
	if n == 0 {
		panic("empty stack")
	}
	return stack.Elements[n-1]
}

func (stack *MinStack) Min() int {
	n := len(stack.MinS)
	return stack.MinS[n-1]
}

func main() {
	fmt.Println("vim-go")
}
