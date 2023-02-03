package main

func gCountIf[T any](arr []T, f func(T) bool) int {
	cnt := 0
	for _, elem := range arr {
		if f(elem) {
			cnt += 1
		}
	}
	return cnt
}
