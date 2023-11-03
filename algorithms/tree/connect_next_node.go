package tree

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description/?envType=daily-question&envId=2023-11-03
type XNode struct {
	Left  *XNode
	Right *XNode
	Next  *XNode
	Val   int
}

func connect(root *XNode) *XNode {
	if root == nil {
		return nil
	}
	q := []*XNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}
