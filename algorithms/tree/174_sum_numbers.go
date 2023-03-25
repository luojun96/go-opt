package main

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *TreeNode) int {
	var dfs func(root *TreeNode, prevSum int) int
	dfs = func(root *TreeNode, prevSum int) int {
		if root == nil {
			return 0
		}
		sum := prevSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}
	return dfs(root, 0)
}

type pairItem struct {
	node *TreeNode
	num  int
}

func sumNumbersByBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	q := []*pairItem{
		{
			node: root,
			num:  root.Val,
		},
	}
	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		left, right, num := item.node.Left, item.node.Right, item.num
		if left == nil && right == nil {
			sum += num
		} else {
			if left != nil {
				q = append(q, &pairItem{left, num*10 + left.Val})
			}
			if right != nil {
				q = append(q, &pairItem{right, num*10 + right.Val})
			}
		}
	}
	return sum
}
