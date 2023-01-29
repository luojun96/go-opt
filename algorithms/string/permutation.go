package main

import "sort"

var a []string

func permutation(s string) []string {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})

	n := len(t)
	perm := make([]byte, 0, n)
	// vis := make([]bool, n)
	var res []string
	var backtrack func(int)

	backtrack = func(i int) {
		if i == n {
			res = append(res, string(perm))
			return
		}

		// for k, v := range vis {
		// 	if v || k > 0 && !vis[k-1] && t[k-1] == t[j] {
		// 		continue
		// 	}
		// 	vis[k] = true
		// 	perm = append(perm, t[j])
		// 	backtrack(i+1)
		// 	perm = perm[:len(perm)-1]
		// 	vis[j] = false
		// }
	}

	backtrack(0)
	return res
}
