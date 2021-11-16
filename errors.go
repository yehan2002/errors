package errors

import "errors"

// Const an error that can be used as a constant.
//
// The error message should be in the form of
// `package: func: error message`
// `package: error message`
//
// Eg: 	`encoding/json: Marshal: unsupported type`
type Const string

// Error an error that can be used as a constant
// Deprecated: use errors.Const for constants and errors.New for other uses.
type Error = Const

var _ error = Const("test")

func (c Const) Error() string { return *(*string)(&c) }

// New returns a new error.
// `e` should be in the form of `package: func: error message` or `package: error message`
func New(e string) error { return errors.New(e) }
