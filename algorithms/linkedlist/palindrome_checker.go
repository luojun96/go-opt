package main

// https://github.com/labuladong/fucking-algorithm/blob/master/%E9%AB%98%E9%A2%91%E9%9D%A2%E8%AF%95%E7%B3%BB%E5%88%97/%E5%88%A4%E6%96%AD%E5%9B%9E%E6%96%87%E9%93%BE%E8%A1%A8.md
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}

	left := head
	right := reverse02(slow)

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse02(head *ListNode) *ListNode {
	prev, cur := head, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// https://leetcode.cn/problems/palindrome-linked-list/
// we cann't use stack to solve the palindrome problem
func _isPalindrome(head *ListNode) bool {
	stack := []int{}
	for head != nil {
		if len(stack) == 0 {
			stack = append(stack, head.Val)
		} else {
			n := len(stack) - 1
			top := stack[n]
			if head.Val != top {
				stack = append(stack, head.Val)
			} else {
				stack = stack[:n]
			}
		}
		head = head.Next
	}
	return len(stack) == 0
	// prev := -1
	// for head != nil {
	// 	if prev > 0 {
	// 		if head.val != prev {
	// 			stack = append(stack, head.Val)
	// 			prev = head.Val
	// 		} else {
	// 			n := len(stack) - 1
	// 			stack = stack[:n]
	// 		}
	// 	} else {
	// 		stack = append(stack, head.Val)
	// 		prev = head.Val
	// 	}

	// 	head = head.Next
	// }
}
