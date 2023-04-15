package tree

import (
	"math"
)

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

		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
