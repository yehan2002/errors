package errors

import (
	"fmt"
	"runtime/debug"
)

// ErrPanic a recovered panic
const ErrPanic Const = "panic"

// Recover recovers from a panic.
// This function must be deferred.
// Usage:
//
//	   func someFunction() (err error){
//			defer Recover(&err)
//	     // do stuff
//	   }
func Recover(err *error) {
	if r := recover(); r != nil {
		recoverErr(err, r, nil)
	}
}

// RecoverStack recovers from a panic and gets the stack trace.
// This function must be deferred.
// Usage:
//
//	   func someFunction() (err error){
//			defer RecoverStack(&err)
//	     // do stuff
//	   }
func RecoverStack(err *error) {
	if r := recover(); r != nil {
		recoverErr(err, r, debug.Stack())
	}
}

func recoverErr(err *error, recovered any, stack []byte) {
	e, ok := recovered.(error)
	if !ok {
		e = fmt.Errorf("%s", recovered)
	}

	if stack != nil {
		e = &panicStackError{err: e, stack: stack}
	}

	// TODO: err may already have a value. keep the original value?
	*err = Cause(ErrPanic, e)
}

type panicStackError struct {
	err   error
	stack []byte
}

func (p *panicStackError) Error() string {
	return fmt.Sprintf("%s\n%s", p.err, p.stack)
}

func (p *panicStackError) Unwrap() error { return p.err }
