package algorithms

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
//
//	input:
//	    3
//	   / \
//	  9  20
//	    /  \
//	   15   7
//	output:
//	[
//	  [3],
//	  [20,9],
//	  [15,7]
//	]
func levelOrderZigzag(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}

	q := []*TreeNode{root}
	level := 1

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

		if level%2 == 0 {
			reverse(arr)
		}

		res = append(res, arr)
		level++
	}

	return res
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l <= r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
