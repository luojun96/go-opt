package main

func hasPathSumByBFS(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	queueNode := []*TreeNode{}
	queueValue := []int{}
	queueNode = append(queueNode, root)
	queueValue = append(queueValue, root.Val)

	for len(queueNode) != 0 {
		curr := queueNode[0]
		queueNode = queueNode[1:]
		temp := queueValue[0]
		queueValue = queueValue[1:]

		if curr.Left == nil && curr.Right == nil {
			if temp == sum {
				return true
			}
			continue

		}
		if curr.Left != nil {
			queueNode = append(queueNode, curr.Left)
			queueValue = append(queueValue, curr.Left.Val+temp)
		}

		if curr.Right != nil {
			queueNode = append(queueNode, curr.Right)
			queueValue = append(queueValue, curr.Right.Val+temp)
		}

	}
	return false
}

func hasPathSumByRecursion(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSumByRecursion(root.Left, sum-root.Val) || hasPathSumByRecursion(root.Right, sum-root.Val)
}
