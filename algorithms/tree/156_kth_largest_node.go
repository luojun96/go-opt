package algorithms

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
// input: root = [5,3,6,2,4,null,null,1], k = 3
//
//	      5
//	     / \
//	    3   6
//	   / \
//	  2   4
//	 /
//	1
//
// output: 4
func kthLargest(root *TreeNode, k int) int {
	var res int
	var inverseInOrder func(node *TreeNode)
	inverseInOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Right != nil {
			inverseInOrder(node.Right)
		}

		k--
		if k == 0 {
			res = node.Val
			return
		}

		if node.Left != nil {
			inverseInOrder(node.Left)
		}
	}

	inverseInOrder(root)
	return res
}
