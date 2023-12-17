package stack

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string-ii/

func removeDuplicateKChars(s string, k int) string {
	var res string
	stack := []byte{}

	var isDuplicate func(arr []byte, val byte) bool
	isDuplicate = func(arr []byte, val byte) bool {
		for _, v := range arr {
			if v != val {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(s); i++ {
		if len(stack) < k-1 {
			stack = append(stack, s[i])
		} else if isDuplicate(stack[len(stack)-k+1:], s[i]) {
			stack = stack[:len(stack)-k+1]
		} else {
			stack = append(stack, s[i])
		}
	}

	res = string(stack)
	return res
}
