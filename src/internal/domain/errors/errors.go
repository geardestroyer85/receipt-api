package errors

import "errors"

var (
	ErrReceiptNotFound       = errors.New("receipt not found")
	ErrInvalidReceiptRequest = errors.New("invalid receipt")
)

type AppError struct {
	Err     error
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(err error, message string, code int) *AppError {
	return &AppError{
		Err:     err,
		Message: message,
		Code:    code,
	}
}
