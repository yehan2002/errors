package errors

type wrappedErr struct {
	err  string
	real error
}

func (w *wrappedErr) Error() string { return w.err + ": " + w.real.Error() }

func (w *wrappedErr) Unwrap() error { return w.real }

// Wrap wraps the given error.
// This does nothing if the error is nil.
func Wrap(str string, err error) error {
	if err == nil {
		return nil
	}
	return &wrappedErr{err: str, real: err}
}
