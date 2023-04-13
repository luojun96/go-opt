package algorithms

import "strconv"

// https://www.nowcoder.com/questionTerminal/7b74386695cc48349af37196f45e62a8
// exp: num = 65, output: 01000001
func convert(num int) string {
	res := strconv.FormatInt(int64(num), 2)
	if len(res) < 8 {
		base := "00000000"
		res = base[:8-len(res)] + res
	}
	return res
}
