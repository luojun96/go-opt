package main

// https://leetcode.cn/problems/insert-into-a-binary-search-tree/solutions/
// input: root = [4,2,7,1,3], val = 5
// output: [4,2,7,1,3,5]
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	p := root
	for p != nil {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				break
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				break
			}
			p = p.Right
		}
	}

	return root
}
