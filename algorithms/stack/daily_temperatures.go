package stack

// https://leetcode.cn/problems/daily-temperatures/

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return nil
	}

	res := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))
	stack = append(stack, 0)

	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] > temperatures[stack[len(stack)-1]] {
			for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
				res[stack[len(stack)-1]] = i - (stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		} else {
			stack = append(stack, i)
		}
	}

	return res
}
