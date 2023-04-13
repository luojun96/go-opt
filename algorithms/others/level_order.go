package algorithms

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil	
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)	
		nodes := []int{}
		for i := 0; i < n; i++ {
			node := queue[0]	
			queue = queue[1:]
			nodes = append(nodes, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, nodes)
	}

	return res
}
