package algorithms

// https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
// input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
// output: true
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

func validate(pushed, popped []int) bool {
	st := []int{}
	j := 0
	for _, e := range pushed {
		st = append(st, e)
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	return len(st) == 0
}
