package main

// https://leetcode.cn/problems/reverse-nodes-in-k-group/
// exp: head = [1, 2, 3, 4, 5], k = 2
// 1 -> 2 -> 3 -> 4 -> 5
// 2 -> 1 -> 4 ->3 -> 5
// 1. reverse the node from head to node k
// 2. take the node (k+1) as head, then recursly invoke reverseKGroup method
// 3. combine the result of two below
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	begin, end := head, head
	for i := 0; i < k; i++ {
		if end == nil {
			return head
		}
		end = end.Next
	}

	newNode := reverse(begin, end)
	begin.Next = reverseKGroup(end, k)
	return newNode
}

func reverse(begin, end *ListNode) *ListNode {
	var prev, curr, next *ListNode
	prev = nil
	curr = begin

	for curr != end {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	
	return prev
}
