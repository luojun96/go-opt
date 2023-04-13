package algorithms

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
// input: s = "abcdefg", k = 2
// output: "cdefgab"
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
