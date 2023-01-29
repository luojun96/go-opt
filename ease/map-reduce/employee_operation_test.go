package main

import (
	"reflect"
	"testing"
)

var list = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

func TestEmployeeCountIf(t *testing.T) {
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})
	if old != 2 {
		t.Errorf("expect %d, got %d\n", 2, old)
	}
}

func TestEmployeeFilterIn(t *testing.T) {
	no_vacation := EmployeeFilterIn(list, func(e *Employee) bool {
		return e.Vacation == 0
	})
	var expected = []Employee{
		{"Hao", 44, 0, 8000},
		{"Jack", 26, 0, 4000},
		{"Marry", 29, 0, 6000},
	}

	if !reflect.DeepEqual(expected, no_vacation) {
		t.Errorf("expect %+v, got %+v\n", expected, no_vacation)

	}
}

func TestEmployeeSumIf(t *testing.T) {
	total_pay := EmployeeSumIf(list, func(e *Employee) int {
		return e.Salary
	})
	var expected = 43500
	if total_pay != expected {
		t.Errorf("expect %d, got %d\n", expected, total_pay)
	}
}
