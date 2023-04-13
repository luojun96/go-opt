package algorithms

// https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
// input: 1->2->4, 1->3->4
// output: 1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	node := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}

	return dummyHead.Next
}
