package tree

// https://leetcode.cn/problems/smallest-missing-genetic-value-in-each-subtree/?envType=daily-question&envId=2023-10-31
func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}

	var dfs func(node int) (map[int]bool, int)
	dfs = func(node int) (map[int]bool, int) {
		genSet := map[int]bool{nums[node]: true}
		for _, child := range children[node] {
			childGenSet, y := dfs(child)
			res[node] = max(res[node], y)
			if len(childGenSet) > len(genSet) {
				genSet, childGenSet = childGenSet, genSet
			}
			for gen := range childGenSet {
				genSet[gen] = true
			}
		}
		for genSet[res[node]] {
			res[node]++
		}
		return genSet, res[node]
	}
	dfs(0)
	return res
}
