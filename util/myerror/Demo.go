package myerror

import "net/http"

func ErrDemoGetDemo(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 10001,
		Message:   "Failed to get demo.",
	}
}
