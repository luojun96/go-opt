package main

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
// input: [1,6,3,2,5]
// output: false
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}

	root := postorder[len(postorder)-1]
	// get the left subtree
	i := 0
	for ; i < len(postorder)-1; i++ {
		if postorder[i] > root {
			break
		}
	}

	// get the right subtree
	j := i
	for ; j < len(postorder)-1; j++ {
		if postorder[j] < root {
			return false
		}
	}

	// check whether the left subtree is BST tree
	left := true
	if i > 0 {
		left = verifyPostorder(postorder[:i])
	}

	// check whether the right subtree is BST tree
	right := true
	if i < len(postorder)-1 {
		right = verifyPostorder(postorder[i : len(postorder)-1])
	}

	return left && right

}
