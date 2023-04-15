package tree

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// output: 3
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}
