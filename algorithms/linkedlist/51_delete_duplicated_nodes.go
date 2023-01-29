package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
// input: head = [1,2,3,3,4,4,5]
// output: [1,2,5]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next
	dup := map[int]bool{}
	
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		} else {
			dup[slow.Val] = true
		}
		fast = fast.Next
	}

	slow.Next = nil
	for head != nil && dup[head.Val] {
		head = head.Next
	}

	node := head
	for node != nil && node.Next != nil {
		if dup[node.Next.Val] {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return head
}
