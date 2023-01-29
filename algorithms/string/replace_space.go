package main

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

func replaceSpace(s string) string {
	if len(s) == 0 {
		return s
	}

	nSpace := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			nSpace++
		}
	}

	newLen := len(s) + nSpace*2
	newStr := make([]byte, newLen)
	oldIndex, newIndex := 0, 0
	for oldIndex < len(s) && newIndex < newLen {
		if s[oldIndex] == ' ' {
			newStr[newIndex] = '%'
			newIndex++
			newStr[newIndex] = '2'
			newIndex++
			newStr[newIndex] = '0'
			newIndex++
			oldIndex++
			// https://go.dev/doc/faq#inc_dec
			// newStr[newIndex++] = '%'
			// newStr[newIndex++] = '2'
			// newStr[newIndex++] = '0'
			// oldIndex++
		} else {
			newStr[newIndex] = s[oldIndex]
			newIndex++
			oldIndex++
		}
	}
	return string(newStr)
}
