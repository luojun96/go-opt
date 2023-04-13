package algorithms

// https://leetcode.cn/problems/subsets/
// nput: nums = [1,2,3]
// output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) [][]int {
	var res [][]int
	set := []int{}
	var dfs func(int)
	dfs = func(curr int) {
		if curr == len(nums) {
			res = append(res, append([]int{}, set...))
			return
		}
		set = append(set, nums[curr])
		dfs(curr + 1)
		set = set[:len(set)-1]
		dfs(curr + 1)
	}
	dfs(0)
	return res
}

