package tree

// https://leetcode.cn/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
func mylowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := mylowestCommonAncestor(root.Left, p, q)
	right := mylowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

func myLowestCommonAncestorByDFS(root, p, q *TreeNode) (res *TreeNode) {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}

		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}

	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}

	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return

}
