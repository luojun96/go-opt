package string

func findTheLongestBalancedSubstring(s string) int {
	res := 0
	count := make([]int, 2)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count[1]++
			res = max(res, 2*min(count[0], count[1]))
		} else if i == 0 || s[i-1] == '1' {
			count[0], count[1] = 1, 0
		} else {
			count[0]++
		}
	}
	return res
}
