package main

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		// iterate the root left and append root into stack until the root == nil
		// then pop the left leef, and append its' value into result
		// and set the root = root.Right
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res

}
