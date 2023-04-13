// https://leetcode.cn/problems/string-matching-in-an-array/
package algorithms

import (
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
func stringMatching(words []string) []string {
	var res []string
	for i, x := range words {
		for j, y := range words {
			if j != i && strings.Contains(y, x) {
				res = append(res, x)
				break
			}
		}
	}
	return res
}
