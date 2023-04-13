package algorithms

// https://leetcode.cn/problems/path-sum/
// input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// output: true
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
