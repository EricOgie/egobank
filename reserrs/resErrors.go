package reserrs

import "net/http"

// MyError for Customed error handling
type MyError struct {
	Code    int    `json:",omitempty" xml:",omitempty"`
	Message string `json:"message" xml:"message"`
}

// ErrorMsg helper function to unset code attribute to
func (err MyError) ErrorMsg() *MyError {
	return &MyError{Message: err.Message}
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
