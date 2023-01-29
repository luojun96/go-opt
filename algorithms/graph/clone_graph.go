package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	return clone(node, map[int]*Node{})
}

func clone(node *Node, visited map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if n, ok := visited[node.Val]; ok {
		return n
	}

	newNode := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	visited[newNode.Val] = newNode

	for _, n := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, clone(n, visited))
	}

	return newNode
}
