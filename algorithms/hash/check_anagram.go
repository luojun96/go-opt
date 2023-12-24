package hash

// https://leetcode.cn/problems/valid-anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]int{}
	for i := 0; i < len(s); i++ {
		m[rune(s[i])]++
	}

	for i := 0; i < len(t); i++ {
		m[rune(t[i])]--
		if m[rune(t[i])] < 0 {
			return false
		}
	}

	return true
}
