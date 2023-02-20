package main

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
// input:
//
//	  3
//	 / \
//	9  20
//	  /  \
//	 15   7
//
// output:
// [
//
//	[3],
//	[9,20],
//	[15,7]
//
// ]
func levelOrderByLevel(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		arr := []int{}
		for size > 0 {
			node := q[0]
			arr = append(arr, node.Val)
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			size--
		}

		res = append(res, arr)
	}

	return res
}
