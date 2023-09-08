package tree

import "math"

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}

	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}

	return minD + 1
}
