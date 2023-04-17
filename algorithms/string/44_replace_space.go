package string

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
// input: s = "We are happy."
// output: "We%20are%20happy."
func replaceSpace(s string) string {
	b := []byte(s)
	res := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, b[i])
		}
	}
	return string(res)
}

func replaceSpaceByStableSlice(s string) string {
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
		} else {
			newStr[newIndex] = s[oldIndex]
			newIndex++
			oldIndex++
		}
	}
	return string(newStr)
}
