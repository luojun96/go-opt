package linkedlist

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
// input: head = [1,2,3,4,5], n = 2
// output: [1,2,3,5]
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func removeNthFromEndByStack(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}

	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

func removeNthFromEndByLength(head *ListNode, n int) *ListNode {
	var getLength func(head *ListNode) int
	getLength = func(head *ListNode) int {
		var length int
		for ; head != nil; head = head.Next {
			length++
		}
		return length
	}

	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}
