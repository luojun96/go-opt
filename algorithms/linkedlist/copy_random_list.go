package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return nil
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}

	if n, ok := cachedNode[node]; ok {
		return n
	}

	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}
