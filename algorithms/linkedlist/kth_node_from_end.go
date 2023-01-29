package main

// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	
	s, f := head, head
	
	for i := 0; i < k - 1; i++ {
		if f.Next != nil {
			f = f.Next	
		} else {
			return nil
		}
	}

	for f.Next != nil {
		f = f.Next
		s = s.Next
	}

	return s
}
