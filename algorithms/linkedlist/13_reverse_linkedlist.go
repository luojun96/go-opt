package main

// https://leetcode.cn/problems/reverse-linked-list/
// input: head = [1,2,3,4,5]
// output: [5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// func reverseList(head *ListNode) *ListNode {
// 	var pre, next *ListNode
//
// 	for head != nil {
// 		next = head.Next
// 		head.Next = pre
// 		pre = head
// 		head = next
// 	}
// 	return pre
// }
