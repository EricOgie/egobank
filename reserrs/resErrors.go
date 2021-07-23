package reserrs

import "net/http"

// MyError for Customed error handling
type MyError struct {
	Code    int
	Message string
}

//NotFoundError helper func
func NotFoundError(msg string) *MyError {
	return &MyError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

// UnexpectedError helper func
func UnexpectedError(msg string) *MyError {
	return &MyError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}
