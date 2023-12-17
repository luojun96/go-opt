package stack

// https://leetcode.cn/problems/valid-parentheses/
// input: s = "()[]{}"
// output: true
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		c := s[i]
		if pairs[c] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

var pairs map[byte]byte = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValidParentheses(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		c := s[i]
		if pair, ok := pairs[c]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
