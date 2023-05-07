package main

import "fmt"

func dvide(a, b int) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in dvide", r)
			err = fmt.Errorf("internal error: %v", r)
		}
	}()

	res = a / b
	return res, err
}

func main() {
	div, err := dvide(5, 0)
	fmt.Println(div, err)
}
