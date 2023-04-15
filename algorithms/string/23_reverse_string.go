package string

// https://leetcode.cn/problems/reverse-string/
// input: s = ["h","e","l","l","o"]
// output: ["o","l","l","e","h"]
func reverseString(s []byte) {
	for left, right := 0, len(s) - 1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}
