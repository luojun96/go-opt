package main

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
// input: s = "We are happy."
// output: "We%20are%20happy."
func replaceSpace(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}
