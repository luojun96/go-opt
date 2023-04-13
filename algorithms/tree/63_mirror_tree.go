package algorithms

// https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/
// input: root = [4,2,7,1,3,6,9]
// output: [4,7,2,9,6,3,1]
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	root.Left = mirrorTree(root.Left)
	root.Right = mirrorTree(root.Right)

	return root
}
