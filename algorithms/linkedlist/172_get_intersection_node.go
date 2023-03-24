package main

// https://leetcode.cn/problems/3u1WK4/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for node := headA; node != nil; node = node.Next {
		vis[node] = true
	}

	for node := headB; node != nil; node = node.Next {
		if vis[node] {
			return node
		}
	}
	return nil
}
