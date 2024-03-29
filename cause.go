package errors

import (
	"errors"
)

type causeError struct{ err, cause error }

func (c *causeError) Error() string { return c.err.Error() + ": " + c.cause.Error() }

func (c *causeError) Unwrap() error { return c.cause }

func (c *causeError) As(target any) bool {
	return errors.As(c.err, target) || errors.As(c.cause, target)
}

func (c *causeError) Is(target error) bool {
	return errors.Is(c.err, target) || errors.Is(c.cause, target)
}

type causeError2 struct{ causeError }

// Cause wraps an error in a another error.
// This returns nil if `cause` is nil.
// If `err` is nil `cause` is returned.
func Cause(err, cause error) error {
	if cause == nil {
		return nil
	}
	if err == nil {
		return cause
	}

	return &causeError2{causeError{err: err, cause: cause}}
}

// CauseStr like `Cause` but cause is a string instead of an error
func CauseStr(err error, cause string) error {
	if err != nil {
		err = Cause(err, New(cause))
	}
	return err
}
