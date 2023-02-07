//go:build go1.20
// +build go1.20

package errors

var _ interface{ Unwrap() []error } = &causeError2{}

func (c *causeError2) Unwrap() []error { return []error{c.cause, c.err} }
