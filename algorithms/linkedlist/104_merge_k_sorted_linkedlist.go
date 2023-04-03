package main

// https://leetcode.cn/problems/merge-k-sorted-lists/
// input: lists = [[1,4,5],[1,3,4],[2,6]]
// output: [1,1,2,3,4,4,5,6]
// if list length > 2, may merge two list, then use the result to merge the third list
func mergeKLists(lists []*ListNode) *ListNode {
	var merge func(list1, list2 *ListNode) *ListNode
	merge = func(list1, list2 *ListNode) *ListNode {
		if list1 == nil {
			return list2
		}
		if list2 == nil {
			return list1
		}
		head := &ListNode{
			Val: 0,
		}
		node := head
		for list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				node.Next = list1
				list1 = list1.Next
			} else {
				node.Next = list2
				list2 = list2.Next
			}
			node = node.Next
		}
		if list1 == nil {
			node.Next = list2
		} else {
			node.Next = list1
		}
		return head.Next
	}
	var merged *ListNode
	for i := 0; i < len(lists); i++ {
		merged = merge(merged, lists[i])
	}
	return merged
}
