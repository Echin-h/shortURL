package errors

import "fmt"

type Error struct {
	Code    int
	Message string
	Origin  string
}

func newError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) WithOrigin(err error) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message,
		Origin:  err.Error(),
	}
}

func (e *Error) WithMessage(message string) *Error {
	return &Error{
		Code:    e.Code,
		Message: fmt.Sprintf(e.Message + " " + message),
	}
}
