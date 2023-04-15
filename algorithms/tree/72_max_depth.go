package tree

// https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/
// input: [3,9,20,null,null,15,7]
//
//	  3
//	 / \
//	9  20
//	  /  \
//	 15   7
//
// output: 3
func maxDepthByRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nLeft := maxDepth(root.Left)
	nRight := maxDepth(root.Right)

	if nLeft > nRight {
		return nLeft + 1
	} else {
		return nRight + 1
	}
}

func maxDepthByBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)

	ans := 0
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		ans++
	}

	return ans
}
