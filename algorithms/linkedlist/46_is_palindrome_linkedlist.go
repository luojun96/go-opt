package main

// https://leetcode.cn/problems/palindrome-linked-list/
// input: head = [1,2,2,1]
// output: true
func check_isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	left := head
	right := reverse_linkedlist(slow)

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse_linkedlist(head *ListNode) *ListNode {
	prev, curr := head, head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
