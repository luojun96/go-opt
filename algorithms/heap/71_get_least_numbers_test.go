package algorithms

import (
	"reflect"
	"testing"
)

func getLeastNumbersByQuickSelectTest(t *testing.T) {
	arr := []int{3, 2, 1}
	k := 2
	res := getLeastNumbersByQuickSelect(arr, k)
	expected := []int{1, 2}
	isEqual := reflect.DeepEqual(res, expected)
	if !isEqual {
		t.Fail()
	}
}
