package main

// https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/
// input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
func buildTreeByPreOrderAndInorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTreeByPreOrderAndInorder(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTreeByPreOrderAndInorder(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}

func buildTreeByIte(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{}
	stack = append(stack, root)

	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}

	return root
}
