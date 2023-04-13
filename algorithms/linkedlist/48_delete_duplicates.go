package algorithms

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
// input: head = [1,1,2]
// output: [1,2]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	} 

	return head
}
