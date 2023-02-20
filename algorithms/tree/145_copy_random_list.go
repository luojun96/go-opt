package main

// https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cachedNode := map[*Node]*Node{}

	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
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

	return deepCopy(head)
}
