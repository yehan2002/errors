package errors

type wrappedError struct {
	err  string
	real error
}

func (w *wrappedError) Error() string { return w.err + ": " + w.real.Error() }

func (w *wrappedError) Unwrap() error { return w.real }

// Wrap wraps the given error with the given context.
// Context should be a string containing more information about where/how the error occurred.
// Example:
//
//	n, err := file.Read(buf)
//	if err != nil {
//		return errors.Wrap("loadConfig: unable to read config", err)
//	}
//
// This returns nil if the error is nil.
func Wrap(ctx string, err error) error {
	if err == nil {
		return nil
	}
	return &wrappedError{err: ctx, real: err}
}
