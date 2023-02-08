package errors

import (
	"errors"
	"testing"
)

func TestRecover(t *testing.T) {
	panicError := Error("a")
	err := panicErr(panicError, false)
	if !errors.Is(err, ErrPanic) {
		t.Fatalf("error is not ErrPanic")
	}
	if !errors.Is(err, panicError) {

	}

	var err2 *wrappedError
	err = panicErr(err2, false)
	if !errors.Is(err, ErrPanic) {
		t.Fatalf("error is not ErrPanic")
	}

}

func panicErr(e error, withStack bool) (err error) {
	if withStack {
		defer RecoverStack(&err)
	} else {
		defer Recover(&err)
	}

	panic(e)
}
