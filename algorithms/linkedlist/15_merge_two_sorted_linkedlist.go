package main

// https://leetcode.cn/problems/merge-two-sorted-lists/
// input: l1 = [1,2,4], l2 = [1,3,4]
// output: [1,1,2,3,4,4]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			node.Next = list2
			list2 = list2.Next
		} else {
			node.Next = list1
			list1 = list1.Next
		}
		node = node.Next
	}

	if list1 != nil {
		node.Next = list1
	}

	if list2 != nil {
		node.Next = list2
	}

	return head.Next
}
