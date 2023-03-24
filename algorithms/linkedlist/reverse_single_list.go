package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
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
