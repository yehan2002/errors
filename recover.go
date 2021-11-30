package errors

import (
	"fmt"
)

// ErrPanic a recovered panic
const ErrPanic Error = "panic"

// Recover recovers from a panic.
// This function must be deferred.
// Usage:
//    func someFunction() (err error){
//		defer Recover(&err)
//      // do stuff
//    }
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
