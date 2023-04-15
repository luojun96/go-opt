package tree

// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
// input: A = [1,2,3], B = [3,1]
// output: false
func isSubStructure(a *TreeNode, b *TreeNode) bool {
	res := false
	if a != nil && b != nil {
		if a.Val == b.Val {
			res = isSubTree(a, b)
		}

		if !res {
			res = isSubStructure(a.Left, b)
		}

		if !res {
			res = isSubStructure(a.Right, b)
		}
	}
	return res

}

func isSubTree(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}

	if a == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return isSubTree(a.Left, b.Left) && isSubTree(a.Right, b.Right)
}
