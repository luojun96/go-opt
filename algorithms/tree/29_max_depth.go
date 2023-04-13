package algorithms

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// input: [3,9,20,null,null,15,7]
// output: 3
func maxDepth(root *TreeNode) int {
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
