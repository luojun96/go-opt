package main

import "math"

func verifyPostorder1(postorder []int) bool {
	stack := make([]int, 0)
	root := math.MaxInt
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}

		for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, postorder[i])
	}

	return true
}
