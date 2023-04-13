package algorithms

func inorderSucessor(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode
	if p.Right != nil {
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}

	node := root
	for node != nil {
		if node.Val > p.Val {
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}

func inorderSuccessorByIteration(root *TreeNode, p *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	var prev, curr *TreeNode
	prev, curr = nil, root
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)	
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev == p {
			return curr	
		}
		prev = curr
		curr = curr.Right
	}
	return nil
}
