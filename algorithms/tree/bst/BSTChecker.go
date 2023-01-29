package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	leftNode := TreeNode{Val: 1, Left: nil, Right: nil}
	rightNode := TreeNode{Val: 3, Left: nil, Right: nil}
	root := TreeNode{
		Val:   2,
		Left:  &leftNode,
		Right: &rightNode,
	}
	res := isValidBST(&root)
	fmt.Println(res)
}

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
