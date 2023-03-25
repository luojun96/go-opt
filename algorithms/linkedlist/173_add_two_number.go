package main

// https://leetcode.cn/problems/lMSNwu/
// input: 2->4->3, 5->6->4
// output: 7->0->8
// explain: 342 + 465 = 807
func addTwoNumbersInLinkedList(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := []int{}
	stack2 := []int{}
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	carry := 0
	var ans *ListNode
	for len(stack1) > 0 || len(stack2) > 0 || carry != 0 {
		a, b := 0, 0
		if len(stack1) > 0 {
			a = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			b = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		cur := a + b + carry
		carry = cur / 10
		cur %= 10
		curNode := &ListNode{Val: cur}
		curNode.Next = ans
		ans = curNode
	}
	return ans
}
