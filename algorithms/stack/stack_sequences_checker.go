package stack

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) > 0 && len(popped) > 0 {
		length := len(pushed)
		stack := make([]int, 0, length)
		i, j := 0, 0
		for j = 0; j < len(popped); j++ {
			if len(stack) > 0 && popped[j] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				for ; i < len(pushed) && pushed[i] != popped[j]; i++ {
					stack = append(stack, pushed[i])
				}
				i++
				if i > len(pushed) {
					return false
				}
			}

		}
	}
	return true
}
