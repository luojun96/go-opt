package main

func Add(x, y int) *int {
	res := 0
	res = x + y
	return &res
}

// build with go build -gcflags "-m -m -l" main.go
func main() {
	Add(1, 2)
}
