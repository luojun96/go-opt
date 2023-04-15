package backtrack

// https://leetcode.cn/problems/combinations/
// input: n = 4, k = 2
// output: [ [1,2], [1,3] [1,4] [2,3] [2,4] [3,4] ]

var res1 [][]int

func combine(n int, k int) [][]int {
	if k <= 0 || n <= 0 {
		return nil
	}
	res1 = [][]int{}
	track := make([]int, 0, k)
	backtrack01(n, k, 1, track)
	return res1
}

func backtrack01(n int, k int, start int, track []int) {
	if k == len(track) {
		res1 = append(res1, track)
		return
	}

	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack01(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}
