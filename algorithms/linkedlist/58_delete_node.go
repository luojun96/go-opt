package linkedlist

// https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// input: head = [4,5,1,9], val = 5
// output: [4,1,9]
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	prev, curr := head, head.Next
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			break
		}
		prev = curr
		curr = curr.Next
	}

	return head
}

func deleteListNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{
		Next: head,
	}

	curr := dummyHead
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			break
		} else {
			curr = curr.Next
		}
	}

	return dummyHead.Next
}
