package string

import (
	"math"
)

// https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/// input: "42"
// output: 42
func strToInt(str string) int {
	state := "start"
	table := map[string][]string{
		"start":  []string{"start", "signed", "number", "end"},
		"signed": []string{"end", "end", "number", "end"},
		"number": []string{"end", "end", "number", "end"},
		"end":    []string{"end", "end", "end", "end"},
	}
	sign := 1
	number := 0
	var getIdx func(c byte) int
	getIdx = func(c byte) int {
		switch c {
		case ' ':
			return 0
		case '-', '+':
			return 1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return 2
		default:
			return 3
		}
	}

	var getMin func(x, y int) int
	getMin = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	for i, _ := range str {
		state = table[state][getIdx(str[i])]
		if state == "number" {
			number = number*10 + (int(str[i]) - '0')
			if sign == 1 {
				number = getMin(number, math.MaxInt32)
			} else if sign == -1 {
				number = getMin(number, -math.MinInt32)
			}
		} else if state == "signed" {
			if str[i] == '-' {
				sign = -1
			}
		}
	}
	return sign * number
}
