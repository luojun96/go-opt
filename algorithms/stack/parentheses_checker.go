package main

// https://leetcode.cn/problems/valid-parentheses/
// string s only contains '('，')'，'{'，'}'，'['，']'
// exp: s = {[]}, output: true
//
//	s = (], output: false
//
// use a hash map to store these three parentheses
// loop the string, store the char into a stack, and pop out the char if the char is right one. After loop the string, the length of stack should be zero
func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if len(stack) == 0 {
			stack = append(stack, c)
		} else {
			top := stack[len(stack)-1]
			if m[c] == top {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, c)
			}
		}
	}

	return len(stack) == 0
}
