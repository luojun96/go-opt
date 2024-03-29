package string

// https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
// func lengthOfLongestSubstring(s string) int {
// 	var max func(i, j int) int
// 	max = func(i, j int) int {
// 		if i > j {
// 			return i
// 		}
// 		return j
// 	}
//
// 	m := map[byte]int{}
// 	n := len(s)
// 	rk, ans := -1, 0
// 	for i := 0; i < n; i++ {
// 		if i != 0 {
// 			delete(m, s[i-1])
// 		}
// 		for rk+1 < n && m[s[rk+1]] == 0 {
// 			m[s[rk+1]]++
// 			rk++
// 		}
// 		ans = max(ans, rk-i+1)
// 	}
// 	return ans
// }

// https://leetcode.cn/problems/longest-substring-without-repeating-characters
// input: "abcabcbb"
// output: 3

func lengthOfLongestSubstring1(s string) int {
	res := 0
	m := map[byte]int{}
	n := len(s)
	rk := -1
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		res = max(res, rk-i+1)
	}
	return res
}
