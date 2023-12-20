package tree

func levelOrderFromBottom(root *TreeNode) [][]int {
	levelOrder := [][]int{}
	if root == nil {
		return nil
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelOrder = append(levelOrder, level)
	}

	for i := 0; i < len(levelOrder)/2; i++ {
		levelOrder[i], levelOrder[len(levelOrder)-1-i] = levelOrder[len(levelOrder)-1-i], levelOrder[i]
	}
	return levelOrder
}

func levelOrderButtom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		a := []int{}
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			a = append(a, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, a)
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
