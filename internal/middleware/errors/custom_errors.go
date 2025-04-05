package errors

import "net/http"

type AppError struct {
	StatusCode int
	Code       string
	Message    string
}

func (e AppError) Error() string {
	return e.Message
}

// NewAppError creates a new AppError with custom status code, error code and message
func NewAppError(statusCode int, code string, message string) AppError {
	return AppError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

func NotFoundError(message string) AppError {
	return NewAppError(http.StatusNotFound, "NOT_FOUND", message)
}

func BadRequestError(message string) AppError {
	return NewAppError(http.StatusBadRequest, "BAD_REQUEST", message)
}

func UnauthorizedError(message string) AppError {
	return NewAppError(http.StatusUnauthorized, "UNAUTHORIZED", message)
}

func ForbiddenError(message string) AppError {
	return NewAppError(http.StatusForbidden, "FORBIDDEN", message)
}

func InternalServerError(message string) AppError {
	return NewAppError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", message)
}

func ConflictError(message string) AppError {
	return NewAppError(http.StatusConflict, "CONFLICT", message)
}
