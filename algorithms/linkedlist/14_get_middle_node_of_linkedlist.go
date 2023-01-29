package main

// https://leetcode.cn/problems/middle-of-the-linked-list/
// input: [1,2,3,4,5]
// output: 3
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
