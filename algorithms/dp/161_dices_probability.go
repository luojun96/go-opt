package dp

// https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof/
// input: 2
// output: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]
func dicesProbability(n int) []float64 {
	res := make([]float64, 6)
	for i := range res {
		res[i] = 1.0 / 6.0
	}
	for i := 2; i <= n; i++ {
		tmp := make([]float64, 5*i+1)
		for j := 0; j < len(res); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += res[j] / 6.0
			}
		}
		res = tmp
	}
	return res
}
