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
		e, ok := r.(error)
		if !ok {
			e = fmt.Errorf("%s", r)
		}

		// `e` will be nil in the following case:
		// 		var err error
		// 		panic(err)
		if e == nil {
			e = New("error(nil)")
		}

		// TODO: err may already have a value. keep the original value?

		*err = Cause(ErrPanic, e)
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
		e, ok := r.(error)
		if !ok {
			e = fmt.Errorf("%s", r)
		}

		// `e` will be nil in the following case:
		// 		var err error
		// 		panic(err)
		if e == nil {
			e = New("error(nil)")
		}

		// TODO: err may already have a value. keep the original value?

		*err = Cause(ErrPanic, &panicStackError{err: e, stack: debug.Stack()})
	}
}

type panicStackError struct {
	err   error
	stack []byte
}

func (p *panicStackError) Error() string {
	return fmt.Sprintf("%s\n%s", p.err, p.stack)
}

func (p *panicStackError) Unwrap() error { return p.err }
