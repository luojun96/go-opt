package dp

import "bytes"

// https://leetcode.cn/problems/generate-parentheses
func generateParenthesis(n int) []string {
	var res []string
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, balance int) {
		if len(path) == n {
			s := bytes.Repeat([]byte{')'}, n*2)
			for _, j := range path {
				s[j] = '('
			}
			res = append(res, string(s))
			return
		}
		for c := 0; c <= balance; c++ {
			path = append(path, i+c)
			dfs(i+c+1, balance-c+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}
