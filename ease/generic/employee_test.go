package main

import (
	"fmt"
	"testing"
)

func TestEmployee(t *testing.T) {
	totalPay := gReduce(employees, 0.0, func(result float32, e Employee) float32 {
		return result + e.Salary
	})
	fmt.Printf("Total Salary: %0.2f\n", totalPay)

	noVacation := gFilterIn(employees, func(e Employee) bool {
		return e.Vacation == 0
	})
	print(noVacation)
}

func TestGSum(t *testing.T) {
	youngerPay := gSum(employees, func(e Employee) float32 {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	fmt.Printf("Total Salary of Young People: %0.2f\n", youngerPay)

	totalVacation := gSum(employees, func(e Employee) int {
		return e.Vacation
	})
	fmt.Printf("Total Vacation: %d day(s)\n", totalVacation)
}

func TestGCountIf(t *testing.T) {
	old := gCountIf(employees, func(e Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people(>40): %d\n", old)

	highPay := gCountIf(employees, func(e Employee) bool {
		return e.Salary >= 6000
	})
	fmt.Printf("High Salary people(>6k): %d\n", highPay)
}
