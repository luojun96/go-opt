package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("Users/luojun/git/golang/ease/slice")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir: =>", string(dir2))

	// when dir1 := path[:sepIndex]
	// dir1 => Users
	// dir2 => luojun/git/golang/ease/slice
	// dir1 => Userssuffix
	// dir2 => uffixn/git/golang/ease/slice

	// when dir1 := path[:sepIndex:sepIndex]
	//	dir1 => Users
	//	dir2 => luojun/git/golang/ease/slice
	//	dir1 => Userssuffix
	//	dir: => luojun/git/golang/ease/slice
}
