package array

// https://leetcode.cn/problems/count-the-number-of-vowel-strings-in-range/?envType=daily-question&envId=2023-11-07

var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func vowelStrings(words []string, left int, right int) int {
	cnt := 0
	for i := left; i <= right; i++ {
		word := words[i]
		fchar, lchar := word[0], word[len(word)-1]
		if vowels[fchar] && vowels[lchar] {
			cnt++
		}
	}
	return cnt
}
