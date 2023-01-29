package main

import "math"

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Val < inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

func isValidBSTByRecursion(root *TreeNode, lower, upper int) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return isValid(root.Left, lower, lower, root.Val) && isValid(root.Right, root.Val, upper)
}
