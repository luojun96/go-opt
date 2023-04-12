package main

// https://leetcode.cn/problems/swap-nodes-in-pairs/
// input: head = [1, 2, 3, 4]
// output: [2, 1, 4, 3]
// use recursion to solve this problem
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func swapPairsByIteration(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	endNode := head.Next.Next
	newHead := swap(head, head.Next)
	head.Next = swapPairs(endNode)
	return newHead
}

func swap(first, second *ListNode) *ListNode {
	second.Next = first
	first.Next = nil
	return second
}
