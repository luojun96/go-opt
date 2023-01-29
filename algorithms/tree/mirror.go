package main

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
