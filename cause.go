package errors

import (
	"reflect"
)

type causeError struct{ err, cause error }

func (c *causeError) Error() string { return c.err.Error() + ": " + c.cause.Error() }

func (c *causeError) Unwrap() error { return c.cause }

func (c *causeError) Is(err error) bool {
	if err == nil {
		return err == c.err
	}
	typ := reflect.TypeOf
	if errV, target := typ(err), typ(c.err); errV.Comparable() && target.Comparable() {
		return err == c.err
	}
	return false
}

//Cause wrap an error in a another error.
//returns nil if `cause` is nil.
//if `err` is nil `cause` is returned.
func Cause(err, cause error) error {
	if cause == nil {
		return nil
	}
	if err == nil {
		return cause
	}
	return &causeError{err: err, cause: cause}
}
