package main

// https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// input: head = [4,5,1,9], val = 5
// output: [4,1,9]

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	pre, cur := head, head.Next

	for cur != nil && cur.Val != val {
		pre = cur
		cur = cur.Next
	}

	if cur != nil {
		pre.Next = cur.Next
	}

	return head
}
