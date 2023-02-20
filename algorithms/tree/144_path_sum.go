package main

// https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
// input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 5
// output: [[5,4,11,2],[5,8,4,5]]
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}

	var dfs func(node *TreeNode, visited []int, sum int)

	dfs = func(node *TreeNode, visited []int, sum int) {
		visited = append(visited, node.Val)
		sum += node.Val
		if node.Left == nil && node.Right == nil && sum == target {
			res = append(res, append([]int{}, visited...))
			return
		}

		if node.Left != nil || node.Right != nil {
			if sum >= target {
				return
			}
			if node.Left != nil {
				dfs(node.Left, visited, sum)
			}
			if node.Right != nil {
				dfs(node.Right, visited, sum)
			}
		}
	}

	dfs(root, []int{root.Val}, root.Val)

	return res
}

func pathSumOfficial(root *TreeNode, target int) (ans []int) {
	path := []int{}
	var dfs func(node *TreeNode, left int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right && left == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, target)
	return
}

type pair struct {
	node *TreeNode
	left int
}

func pathSumByBFS(root *TreeNode, target int) (ans [][]int) {
	if root == nil {
		return
	}

	parent := map[*TreeNode]*TreeNode{}

	getPath := func(node *TreeNode) (path []int) {
		for ; node != nil; node = parent[node] {
			path = append(path, node.Val)
		}
		for i, j := 0, len(path)-1; i < j; i++ {
			path[i], path[j] = path[j], path[i]
			j--
		}
		return
	}

	queue := []pair{
		{root, target},
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		node := p.node
		left := p.left - node.Val
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				ans = append(ans, getPath(node))
			}
		} else {
			if node.Left != nil {
				parent[node.Left] = node
				queue = append(queue, pair{node.Left, left})
			}
			if node.Right != nil {
				parent[node.Right] = node
				queue = append(queue, pair{node.Right, left})
			}
		}
	}

	return

}
