package stack

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/

func removeDuplicates(s string) string {
	var res string
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	res = string(stack)
	return res
}
