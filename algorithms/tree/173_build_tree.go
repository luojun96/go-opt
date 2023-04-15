package tree

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTreeByInorderAndPostorder(inorder []int, postorder []int) *TreeNode {
	m := map[int]int{}
	for i, v := range inorder {
		m[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		inordrRootIndex := m[val]
		root.Right = build(inordrRootIndex+1, right)
		root.Left = build(left, inordrRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}
