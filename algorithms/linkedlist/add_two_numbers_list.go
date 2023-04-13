package algorithms

// https://leetcode.cn/problems/add-two-numbers-ii/
// input: l1 = [7,2,4,3], l2 = [5,6,4]
// output: [7,8,0,7]
func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)

	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	res := &ListNode{}
	carry := 0
	for len(stack1) > 0 || len(stack2) > 0 || carry != 0 {
		var a, b int
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
		curNode.Next = res
		res = curNode
	}

	return res
}
