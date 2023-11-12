package graph

// https://leetcode.cn/problems/couples-holding-hands/description/?envType=daily-question&envId=2023-11-11

func minSwapsCouples(row []int) int {
	res := 0
	n := len(row)
	g := make([][]int, n/2)
	for i := 0; i < n; i += 2 {
		l, r := row[i]/2, row[i+1]/2
		if l != r {
			g[l] = append(g[l], r)
			g[r] = append(g[r], l)
		}
	}
	vis := make([]bool, n/2)
	for i, vs := range vis {
		if !vs {
			vis[i] = true
			cnt := 0
			q := []int{i}
			for len(q) > 0 {
				res++
				v := q[0]
				q = q[1:]
				for _, w := range g[v] {
					if !vis[w] {
						vis[w] = true
						q = append(q, w)
					}
				}
			}
			res += cnt - 1
		}
	}
	return res
}
