package main

import (
	"fmt"
	"reflect"
)

type RefContainer struct {
	s reflect.Value
}

func NewRefContainer(t reflect.Type, size int) *RefContainer {
	if size <= 0 {
		size = 64
	}
	return &RefContainer{
		s: reflect.MakeSlice(reflect.SliceOf(t), 0, size),
	}
}

func (c *RefContainer) Put(val interface{}) error {
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		return fmt.Errorf("Put: cannot put a %T into a slice of %s", val, c.s.Type().Elem())
	}
	c.s = reflect.Append(c.s, reflect.ValueOf(val))
	return nil
}

func (c *RefContainer) Get(refval interface{}) error {
	if reflect.ValueOf(refval).Kind() != reflect.Ptr || reflect.ValueOf(refval).Elem().Type() != c.s.Type().Elem() {
		return fmt.Errorf("Get: needs *%s but got %T", c.s.Type().Elem(), refval)
	}
	reflect.ValueOf(refval).Elem().Set(c.s.Index(0))
	c.s = c.s.Slice(1, c.s.Len())
	return nil
}
