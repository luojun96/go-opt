package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type authError struct {
	opr string
	err error
}

func (e *authError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.opr, e.err)
}

type causer interface {
	Cause() error
}

func (e *authError) Cause() error {
	return e.err
}

func handleError() error {
	var err error

	if err != nil {
		return errors.Wrap(err, "read failed")
	}

	switch err := errors.Cause(err).(type) {
	case *authError:
		// handle specifically
		return errors.Wrap(err, "authorization error")
	default:
		// unknown error
		return errors.Wrap(err, "unknown error")
	}
}
