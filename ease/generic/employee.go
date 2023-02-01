package main

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   float32
}

var employees = []Employee{
	{"Hao", 44, 0, 8000.5},
	{"Bob", 34, 10, 5000.5},
	{"Alice", 23, 5, 9000.0},
	{"Jack", 26, 0, 4000.0},
	{"Tom", 48, 9, 7500.5},
	{"Marry", 29, 0, 6000.0},
	{"Mike", 32, 8, 4000.0},
}
