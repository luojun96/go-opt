package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type authorizationError struct {
	operation string
	err       error
}

func (e *authorizationError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}

type causer interface {
	Cause() error
}

func (e *authorizationError) Cause() error {
	return e.err
}

func handleError() error {
	var err error

	if err != nil {
		return errors.Wrap(err, "read failed")
	}

	switch err := errors.Cause(err).(type) {
	case *authorizationError:
		// handle specifically
		return errors.Wrap(err, "authorization error")
	default:
		// unknown error
		return errors.Wrap(err, "unknown error")
	}
}
