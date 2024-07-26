package error

import "fmt"

type AppError struct {
	BaseError error
	Message   string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%v: %+v", e.Message, e.BaseError)
}
func NewAppError(message string, baseError error) *AppError {
	return &AppError{
		BaseError: baseError,
		Message:   message,
	}
}
