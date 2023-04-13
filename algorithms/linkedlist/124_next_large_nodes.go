package algorithms

// https://leetcode.cn/problems/next-greater-node-in-linked-list/
// input: [2,1,5]
// output: [5,5,0]
func nextLargerNodes(head *ListNode) []int {
	st := []int{0}
	ans := []int{}
	var dfs func(node *ListNode)
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}
		dfs(node.Next)
		for len(st) > 1 && st[len(st)-1] <= node.Val {
			st = st[:len(st)-1]
		}
		ans = append(ans, st[len(st)-1])
		st = append(st, node.Val)
	}
	dfs(head)
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return ans
}
