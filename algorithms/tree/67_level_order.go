package main

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// input: [3,9,20,null,null,15,7]
// output: [3,9,20,15,7]
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return res
}
