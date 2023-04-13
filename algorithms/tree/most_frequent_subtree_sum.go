package algorithms

// https://leetcode.cn/problems/most-frequent-subtree-sum/
// input: root = [5,2,-5]
// output: 2
func findFrequentTreeSum(root *TreeNode) []int {
	m := map[*TreeNode]int{}

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if v, ok := m[node]; ok {
			return v
		}

		var left, right int
		if node.Left != nil {
			left = dfs(node.Left)
		}
		if node.Right != nil {
			right = dfs(node.Right)
		}

		value := node.Val + left + right
		m[node] = value
		return value
	}
	dfs(root)

	maxCount := 0
	res := []int{}
	cache := map[int]int{}
	for _, v := range m {
		count := cache[v]
		count++
		cache[v] = count
		if maxCount < count {
			maxCount = count
			res = []int{v}
		} else if maxCount == count {
			res = append(res, v)
		}
	}
	return res
}

func findFrequentTreeSumByDFS(root *TreeNode) []int {
	cnt := map[int]int{}
	maxCnt := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		cnt[sum]++
		if cnt[sum] > maxCnt {
			maxCnt = cnt[sum]
		}
		return sum
	}
	dfs(root)

	res := []int{}
	for s, c := range cnt {
		if c == maxCnt {
			res = append(res, s)
		}
	}
	return res
}
