package main

import (
	"fmt"
	"strings"
)

func execMapStr() {
	var list = []string{"Luo", "Jun", "Home"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)
}

func execMapInt() {
	var list = []string{"go", "shell", "js", "docker", "kubernetes"}

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", y)
}
