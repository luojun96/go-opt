package tree

import "sort"

// https://leetcode.cn/problems/count-complete-tree-nodes/
// input: root = [1,2,3,4,5,6]
// output: 6
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}

	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}
