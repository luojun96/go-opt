package main

import "fmt"

func execReduce() {
	var list = []string{"go", "shell", "js", "docker", "kubernetes"}

	x := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", x)
}
