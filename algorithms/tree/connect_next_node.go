package tree

import "fmt"

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
	printTree(root)
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
	fmt.Println("after setting next node:")
	printTree(root)
	return root
}

// print tree by preorder
func printTree(root *XNode) {
	if root == nil {
		return
	}
	fmt.Printf("node: %+v\n", root)
	printTree(root.Left)
	printTree(root.Right)
}
