package tree

// https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof
func treeToDoublylist(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var prev, head *TreeNode
	var convert func(node *TreeNode)
	convert = func(node *TreeNode) {
		if node == nil {
			return
		}
		convert(node.Left)
		if prev == nil {
			head = node
		} else {
			prev.Right = node
			node.Left = prev
		}
		prev = node
		convert(node.Right)
	}
	convert(root)
	return head
}
