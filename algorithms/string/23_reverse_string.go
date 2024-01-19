package string

// https://leetcode.cn/problems/reverse-string/
// input: s = ["h","e","l","l","o"]
// output: ["o","l","l","e","h"]
func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l++ {
		s[l], s[r] = s[r], s[l]
		r--
	}
}
