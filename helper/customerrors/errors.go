package customerror

import (
	"errors"
	"net/http"
)

// Base error types
var (
	ErrBadRequest = errors.New("bad request")
	ErrConflict   = errors.New("conflict")
	ErrNotFound   = errors.New("not found")
)

// CustomError represents an error with additional context
type CustomError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"error"`
}

// Error implementation for the error interface
func (e *CustomError) Error() string {
	return e.Message
}

// NewBadRequest creates a new bad request error with a custom message
func NewBadRequest(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

// NewConflict creates a new conflict error with a custom message
func NewConflict(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}

// NewNotFound creates a new not found error with a custom message
func NewNotFound(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

// GetStatusCode returns the HTTP status code for an error
func GetStatusCode(err error) int {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.StatusCode
	}

	if errors.Is(err, ErrBadRequest) {
		return http.StatusBadRequest
	} else if errors.Is(err, ErrConflict) {
		return http.StatusConflict
	} else if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound
	} else if errors.Is(err, ErrValidation) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
