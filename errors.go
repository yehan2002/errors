package errors

//Error an error that can be used as a constant
type Error string

func (e *Error) Error() string { return *(*string)(e) }
