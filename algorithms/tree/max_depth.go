package algorithms

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nLeft := MaxDepth(root.Left)
	nRight := MaxDepth(root.Right)

	if nLeft > nRight {
		return nLeft + 1
	} else {
		return nRight + 1
	}
}
