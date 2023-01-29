//go:build go1.20
// +build go1.20

package errors

func (c *causeError2) Unwrap() []error { return []error{c.cause, c.err} }
