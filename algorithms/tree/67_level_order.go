package main

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// input: [3,9,20,null,null,15,7]
// output: [3,9,20,15,7]
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			node := q[0]
			res = append(res, node.Val)
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			size--
		}
	}

	return res
}
