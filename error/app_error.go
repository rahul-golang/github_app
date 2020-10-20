package error

import (
	"net/http"
)

const (
	BadRequestErrorType ErrType = iota + 1
	NotFoundErrorType
	InternalServerErrorType
)

const (
	BadRequestErrorCode     = "ERR_APP_BAD_REQUEST_ERROR"
	NotFoundErrorCode       = "ERR_APP_NOT_FOUND_ERROR"
	InternalServerErrorCode = "ERR_APP_INTERNAL_SERVER_ERROR"
	NotAuthorizedErrorCode  = "ERR_APP_NOT_AUTHORIZED_ERROR"
)

type ErrType uint

func (errorType ErrType) New(errorCode ErrCode, errorMessage string, statusCode int) *APPError {
	return &APPError{
		ErrType: errorType,
		ErrResponse: ErrResponse{
			ErrCode:    errorCode,
			ErrMessage: errorMessage,
		},
		StatusCode: statusCode,
	}
}

type ErrCode string

type ErrResponse struct {
	ErrCode    ErrCode `json:"error_code"`
	ErrMessage string  `json:"error_message"`
}

type APPError struct {
	StatusCode  int
	ErrType     ErrType
	ErrResponse ErrResponse
}

var ErrRecordNotFound = NotFoundErrorType.New(NotFoundErrorCode, "record not found", http.StatusNotFound)
var ErrNotAuthorized = BadRequestErrorType.New(NotAuthorizedErrorCode, "not authorized to perform this action", http.StatusUnauthorized)

func BadRequestErrorFunc(errorMessage string) *APPError {
	return BadRequestErrorType.New(BadRequestErrorCode, errorMessage, http.StatusBadRequest)
}

func InternalServerErrorFunc(errorMessage string) *APPError {
	return InternalServerErrorType.New(InternalServerErrorCode, errorMessage, http.StatusInternalServerError)
}
