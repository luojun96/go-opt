package algorithms

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)

		if node.Left != nil {
			queue = append(res, node.Left)
		}

		if node.Right != nil {
			queue = append(res, node.Right)
		}
	}

	return res
}
