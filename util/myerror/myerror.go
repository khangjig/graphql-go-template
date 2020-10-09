package myerror

import "github.com/pkg/errors"

type MyError struct {
	Raw       error
	ErrorCode int
	HTTPCode  int
	Message   string
}

func (e MyError) Error() string {
	if e.Raw != nil {
		return errors.Wrap(e.Raw, e.Message).Error()
	}

	return e.Message
}

func NewError(err error, httpCode, errCode int, message string) MyError {
	return MyError{
		Raw:       err,
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   message,
	}
}
