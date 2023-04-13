package algorithms

import (
	"math"
)

// https://leetcode.cn/problems/reverse-integer/
// x = 123  => 321
// x = -123 => -321
// x = 120 => 21
// x = 0 => 0
// x = 123, 123 % 10 = 3, 123 / 10 = 12; 12 % 10 = 2, 12 / 10 = 1; 1 % 10 = 1, 1 / 10 = 0 ----
// leading zero case: x = 120; 120 % 10 = 0, 120 / 10 = 12; 12 % 10 = 2, 12 / 10 = 1; 1 % 10 = 1, 1 / 10 = 0 ---- 021 -> 21, if the last number is 0, then don't add
// navigate number, if x < 0, then navigate = true
func reverse(x int) int {
	res := 0
	for x != 0 {
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		res = res*10 + digit
	}
	return res
	// fmt.Printf("max int: %d", math.MaxInt32)
	// navigate := false
	// s := ""
	// if x < 0 {
	// 	navigate = true
	// 	x = -x
	// }
	// for x/10 != 0 {
	// 	s += fmt.Sprint(x % 10)
	// 	x = x / 10
	// }

	// fmt.Printf("reversed number without first postion: %s\n", s)
	// if x/10 == 0 {
	// 	s += fmt.Sprint(x % 10)
	// }

	// fmt.Printf("reversed number: %s\n", s)

	// res, _ := strconv.ParseInt(s, 10, 0)

	// if res > math.MaxInt32 && res < math.MinInt32 {
	// 	return 0
	// }

	// if navigate {
	// 	return -1 * int(res)
	// } else {
	// 	return int(res)
	// }
}
