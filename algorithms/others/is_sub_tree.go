package algorithms

type treeNode struct {
	Val int
	Left *treeNode
	Right *treeNode
}

func isSubStructure(a *treeNode, b treeNode) bool {
	res := false
	if a != nil && b != nil {
		if a.Val == b.Val {
			res = isSubTree(a.Left, b)	
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

func isSubTree(a *treeNode, b *treeNode) bool {
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
