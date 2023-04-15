package tree

// https://leetcode.cn/problems/delete-node-in-a-bst/
// 二叉搜索树的节点删除分三种情况处理
// 1. 要删除的节点没有子节点：只需要直接将父节点中，指向要删除节点的指针置为null；
// 2. 要删除的节点只有一个子节点（左子节点/右子节点）：只需要更新父节点中，指向要删除节点的指针，让它指向要删除节点的子节点；
// 3. 要删除的节点有两个子节点：需要找到右子树中的最小节点，把它替换到要删除的节点上，然后再删除这个最小节点，
func deleteNode(root *TreeNode, key int) *TreeNode {
	var node, pNode *TreeNode
	node = root
	pNode = nil
	for node != nil {
		if node.Val == key {
			break
		}
		pNode = node
		if node.Val < key {
			node = node.Right
		} else if node.Val > key {
			node = node.Left
		}
	}
	if node == nil {
		return root
	}

	if node.Left != nil && node.Right != nil {
		minNode := node.Right
		minPNode := node
		for minNode.Left != nil {
			minPNode = minNode
			minNode = minNode.Left
		}
		node.Val = minNode.Val
		node = minNode
		pNode = minPNode
	}

	var child *TreeNode
	if node.Left != nil {
		child = node.Left
	} else if node.Right != nil {
		child = node.Right
	} else {
		child = nil
	}

	if pNode == nil {
		root = child
	} else if pNode.Left == node {
		pNode.Left = child
	} else {
		pNode.Right = child
	}
	return root
}
