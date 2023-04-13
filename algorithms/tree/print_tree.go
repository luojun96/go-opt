package algorithms

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	q := []*TreeNode{root}
	l := 1
	for len(q) > 0 {
		nodes := []int{}
		size := len(q)
		for i := 0; i < size; i++ {
			cNode := q[i]
			node := q[size-i-1]
			nodes = append(nodes, cNode.Val)
			if (l+1)%2 == 0 {
				if node.Right != nil {
					q = append(q, node.Right)
				}
				if node.Left != nil {
					q = append(q, node.Left)
				}
			} else {
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
		}
		if len(q) > size {
			q = q[size:]
		} else {
			q = nil
		}
		res = append(res, nodes)
		l++
	}
	return res
}
