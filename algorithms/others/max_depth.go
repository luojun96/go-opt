package main

func maxDepth(root *treeNode) int {
	if root == nil {
		return 0	
	}

	nLeft := maxDepth(root.Left)
	nRight := maxDepth(root.Right)
	if nLeft > nRight {
		return nLeft + 1	
	} else {
		return nRight + 1	
	}

}
