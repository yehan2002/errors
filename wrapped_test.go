package errors

import (
	"errors"
	"io"
	"testing"
)

func TestWrapped(t *testing.T) {
	wrapMsg := "unexpected error while reading file"
	err := Wrap(wrapMsg, io.ErrUnexpectedEOF)

	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatal("errors.Is returned incorrect result")
	}

	if errors.Is(err, io.ErrClosedPipe) {
		t.Fatal("errors.Is returned incorrect result")
	}

	expectedMsg := wrapMsg + ": " + io.ErrUnexpectedEOF.Error()
	if err.Error() != expectedMsg {
		t.Fatalf("Wrap returned incorrect error message: expected %q got %q", expectedMsg, err.Error())
	}
}

func TestWrappedNil(t *testing.T) {
	err := Wrap("unexpected error", nil)
	if err != nil {
		t.Fatalf("wrapping nil error did not return a nil error")
	}

}
