package algorithms

// https://leetcode.cn/problems/find-bottom-left-tree-value/
// input: root = [2,1,3]
// output: 1
// 广度优先搜索，直到最后一层的第一个节点
func findBottomLeftValue(root *TreeNode) int {
	res := root.Val
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}

func findBottomLeftValueByDFS(root *TreeNode) int {
	res, curHeight := root.Val, 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
		if height > curHeight {
			curHeight = height
			res = node.Val
		}
	}
	dfs(root, 0)
	return res
}

func findBottomLeftValueByBFS(root *TreeNode) int {
	res := root.Val
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		res = node.Val
	}
	return res
}
