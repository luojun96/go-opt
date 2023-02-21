package main

// https://leetcode.cn/problems/symmetric-tree/
// input: root = [1,2,2,3,4,4,3]
// output: true
func isSymmetricTree(root *TreeNode) bool {
	var check func(p, q *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil {
			return false
		}

		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}

	return check(root, root)
}
