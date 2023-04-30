package main

import (
	"testing"
)

func TestPut(t *testing.T) {
	initContainer := &Container{}
	initContainer.Put(7)
	initContainer.Put(42)

	// assert that actual type is int
	elem, ok := initContainer.Get().(int)
	if !ok {
		t.Error("Unable to read int from iniContainer")
	}
	if elem != 7 {
		t.Errorf("expect 7, got %v\n", elem)
	}
}
