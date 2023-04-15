package tree

// https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
// input: [3,9,20,null,null,15,7]
//
//	  3
//	 / \
//	9  20
//	  /  \
//	 15   7
//
// output: true
func isBalanced(root *TreeNode) bool {
	depth := 0
	return balance(root, &depth)
}

func balance(root *TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}

	left, right := 0, 0
	if balance(root.Left, &left) && balance(root.Right, &right) {
		diff := left - right
		if diff <= 1 && diff >= -1 {
			if left > right {
				*depth = 1 + left
			} else {
				*depth = 1 + right
			}
			return true
		}
	}

	return false
}
