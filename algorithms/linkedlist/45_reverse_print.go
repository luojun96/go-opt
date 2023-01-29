package main

// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
// head = [1,3,2]
// [2,3,1]
func reversePrint(head *ListNode) []int {
	node := &ListNode{}
	count := 0
	for head != nil {
		next := head.Next
		head.Next = node
		node = head
		head = next
		count++
	}

	res := make([]int, 0, count)
	for node != nil && node.Next != nil {
		res = append(res, node.Val)
		node = node.Next
	}

	return res
}

func reversePrintByStack(head *ListNode) []int {
	stack := []int{}
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	res := []int{}
	for len(stack) > 0 {
		n := len(stack) - 1
		res = append(res, stack[n])
		stack = stack[:n]
	}
	return res
}
