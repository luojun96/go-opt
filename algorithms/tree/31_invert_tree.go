package tree

// https://leetcode.cn/problems/invert-binary-tree/
// input: root = [4,2,7,1,3,6,9]
// output: [4,7,2,9,6,3,1]
// input:
//
//		    4
//		  /   \
//		 2	   7
//		/ \   / \
//	 1   3 6   9
//
// output:
//
//		    4
//		  /   \
//		 7		 2
//		/ \   / \
//	 9   6 3   1
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func invertBTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertBTree(root.Left)
	right := invertBTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
