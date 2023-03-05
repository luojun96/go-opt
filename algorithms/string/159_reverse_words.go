package main

// https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/
// input: "the sky is blue"
// output: "blue is sky the"
func reverseWords(s string) string {
	trimString(s)

	stack := []string{}
	word := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word = append(word, s[i])
		} else {
			if len(word) > 0 {
				stack = append(stack, string(word))
				word = []byte{}
			}
		}
	}
	if len(word) > 0 {
		stack = append(stack, string(word))
	}

	return reverseSilceToString(stack)
}

func trimString(s string) {
	l, r := 0, len(s)-1
	for ; l <= r && s[l] == ' '; l++ {
	}
	for ; r > l && s[r] == ' '; r-- {
	}
	s = s[l : r+1]
}

func reverseSilceToString(arr []string) string {
	var res string
	for len(arr) > 0 {
		str := arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		if len(arr) == 0 {
			res += str
		} else {
			res += str + " "
		}
	}
	return res
}
