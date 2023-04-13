package algorithms

// https://leetcode.cn/problems/binary-tree-postorder-traversal/
// input: root = [1,null,2,3]
// output: [3,2,1]
func postorderTraversalByRecursion(root *TreeNode) (res []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}

func postorderTraversalByIteration(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}
