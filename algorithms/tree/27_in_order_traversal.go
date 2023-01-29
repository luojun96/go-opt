package main

// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// input: root = [1,null,2,3]
// output: [1,3,2]
func inorderTraversal(root *TreeNode) []int {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			retuen 	
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}	
	inorder(root)
	return
}

func inorderTraversalByIteration(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res= append(res, root.Val)
		root = root.Right
	}
	return
}
