package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.cn/problems/is-subsequence/

func isSubsequence01(s string, t string) bool {
	lt := len(t)
	// index := [256][]int{}
	index := make([][]int, 256)
	for i := 0; i < lt; i++ {
		b := t[i]
		ty := reflect.TypeOf(index[b])
		fmt.Println(ty)
		index[b] = append(index[b], i)
	}

	fmt.Printf("index: %v", index)
	fmt.Println()
	j := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if index[b] == nil {
			return false
		}
		p := getLeftBound(index[b], j)
		if p == -1 {
			return false
		}
		j = index[b][p] + 1
	}
	return true
}

func getLeftBound(a []int, t int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := l + (r-l)/2
		if a[mid] >= t {
			if mid == 0 || a[mid-1] < t {
				return mid
			} else {
				r = mid - 1
			}
		} else {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	res := isSubsequence01("abce", "adbecc")
	fmt.Println(res)
}
