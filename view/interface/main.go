package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v interface{}
	v = (*int)(nil)
	fmt.Printf("isNil: %t\n", IsNil(v))
	fmt.Println(v == nil)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
