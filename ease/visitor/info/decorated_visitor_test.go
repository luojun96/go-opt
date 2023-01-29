package main

import (
	"fmt"
	"testing"
)

func TestDecoratedVisitor(t *testing.T) {
	info := Info{}

	fn := func(info *Info, err error) error {
		info.Name = "Luo Jun"
		info.Namespace = "Developer"
		info.OtherThings = "free"
		return nil
	}

	var v Visitor = &info
	var nameVisitor, logVisitor, otherVisitor VisitorFunc
	nameVisitor = func(i *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(i, err)
		if err == nil {
			fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	}

	logVisitor = func(i *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(i, err)
		fmt.Println("LogVisitor() after call function")
		return err
	}

	otherVisitor = func(i *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(i, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	}
	v = NewDecoratedVisitor(v, nameVisitor, logVisitor, otherVisitor)

	v.Visit(fn)
}
