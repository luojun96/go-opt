package string

// https://leetcode.cn/problems/is-subsequence
// input: s = "abc", t = "ahbgdc"
// output: true
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}
