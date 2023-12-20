package tree

// https://leetcode.cn/problems/binary-tree-right-side-view/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			if i == size-1 {
				res = append(res, node.Val)
			}
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}
