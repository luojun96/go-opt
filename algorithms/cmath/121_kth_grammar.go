package cmath

// https://leetcode.cn/problems/k-th-symbol-in-grammar
// input: n = 2, k = 1
// output: 0
func kthGrammar(n, k int) int {
	if n == 1 {
		return 0
	}
	return k&1 ^ 1 ^ kthGrammar(n-1, (k+1)/2)
}
