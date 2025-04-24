package errors

import "fmt"

// CustomError defines a custom error type that includes a message and a code.
type CustomError struct {
    Code    int
    Message string
}

// New creates a new CustomError with the given code and message.
func New(code int, message string) *CustomError {
    return &CustomError{
        Code:    code,
        Message: message,
    }
}

// Error implements the error interface for CustomError.
func (e *CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// IsNotFound checks if the error is a not found error.
func IsNotFound(err error) bool {
    if customErr, ok := err.(*CustomError); ok {
        return customErr.Code == 404
    }
    return false
}

// IsUnauthorized checks if the error is an unauthorized error.
func IsUnauthorized(err error) bool {
    if customErr, ok := err.(*CustomError); ok {
        return customErr.Code == 401
    }
    return false
}