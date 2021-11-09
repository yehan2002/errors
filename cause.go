package errors

import (
	"errors"
)

type causeError struct{ err, cause error }

func (c *causeError) Error() string { return c.err.Error() + ": " + c.cause.Error() }

func (c *causeError) Unwrap() error { return c.cause }

func (c *causeError) Is(target error) bool { return errors.Is(c.err, target) }

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
	return &causeError{err: err, cause: cause}
}
