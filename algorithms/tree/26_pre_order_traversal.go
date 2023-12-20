package tree

// https://leetcode.cn/problems/binary-tree-preorder-traversal/
// input: root = [1,null,2,3]
// output: [1,2,3]
func preorderTraversal(root *TreeNode) []int {
	var vals []int
	var preorder func(*TreeNode)

	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)
	return vals
}

func preorderTraversalByIteration(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return
}

func preorderByIteration(root *TreeNode) (vals []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{
		root,
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append(vals, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return
}
