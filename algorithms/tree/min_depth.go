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

func minDepthByBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	minDepth := 0
	for len(q) > 0 {
		size := len(q)
		minDepth++
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return minDepth
			}
		}
	}
	return minDepth
}
