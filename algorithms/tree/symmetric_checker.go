package main

// https://leetcode.cn/problems/symmetric-tree/
// 		1
// 	  2   2
//   3 4 4 3
// loop the tree by tier, and check in each tier that: 1. node value; 2. node postion, right node or left node
func isSymmetric(root *TreeNode) bool {
	q := []*TreeNode{root}
	tier := 1
	for len(q) > 0 {
		size := len(q)
		if tier != 1 && size%2 != 0 {
			return false
		}
		l, r := 0, len(q)-1
		for l < r {
			if q[l].Val != q[r].Val {
				return false
			}
			// fmt.Printf("left node %v\n", q[l])
			// fmt.Printf("right node %v\n", q[r])
			if (q[l].Left == nil && q[r].Right != nil) || (q[l].Left != nil && q[r].Right == nil) || (q[l].Left != nil && q[r].Right != nil && q[l].Left.Val != q[r].Right.Val) ||
				(q[l].Right == nil && q[r].Left != nil) || (q[l].Right != nil && q[r].Left == nil) || (q[l].Right != nil && q[r].Left != nil && q[l].Right.Val != q[r].Left.Val) {
				return false
			}
			l++
			r--
		}

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		tier++
	}
	return true
}

func isSymmetricByRecursion(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, p.Right) && check(p.Right, q.Left)
}
