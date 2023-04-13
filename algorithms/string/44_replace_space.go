package algorithms

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
