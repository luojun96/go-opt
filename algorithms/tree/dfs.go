package main

func dfs(node *TreeNode, target *TreeNode, visited map[*TreeNode]bool) bool {
	if node.Val == target.Val {
		return true
	}

	if node.Left != nil && visited[node.Left] == false {
		visited[node.Left] = true
		if dfs(node.Left, target, visited) == true {
			return true
		}
	}

	if node.Right != nil && visited[node.Right] == false {
		visited[node.Right] = true
		if dfs(node.Right, target, visited) == true {
			return true
		}
	}

	return false
}
