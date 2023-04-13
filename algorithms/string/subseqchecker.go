package algorithms

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for ; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	fmt.Println(i)
	return i == len(s)
}

