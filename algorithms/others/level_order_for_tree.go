package others

type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

func levelOrderBTree(root *TNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	queue := []*TNode{}
	if len(queue) > 0 {
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
