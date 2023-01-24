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
//
// Deprecated: use errors.Const for constants and [New] for other uses.
type Error = Const

var _ error = Const("test")

func (c Const) Error() string { return *(*string)(&c) }

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
//
// This function is an alias to [New] in the standard library.
func New(e string) error { return errors.New(e) }
