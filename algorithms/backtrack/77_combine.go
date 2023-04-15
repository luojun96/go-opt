package backtrack

// https://leetcode.cn/problems/combinations/
// input: n = 4, k = 2
// output: [ [1,2], [1,3] [1,4] [2,3] [2,4] [3,4] ]

func combineByDFS(n int, k int) [][]int {
	temp := []int{}
	var res [][]int
	var dfs func(int)
	dfs = func(cur int) {
		if len(temp)+(n-cur+1) < k {
			return
		}

		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			res = append(res, comb)
			return
		}

		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return res
}
