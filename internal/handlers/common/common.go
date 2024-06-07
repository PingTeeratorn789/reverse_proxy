package common

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type ErrorCode int

const (
	ErrorCodeBadRequest             ErrorCode = 4000
	ErrorCodeUnauthorized           ErrorCode = 4001
	ErrorCodePageNotFound           ErrorCode = 4004
	ErrorCodeStatusMethodNotAllowed ErrorCode = 4005
	ErrorCodeRequestEntityTooLarge  ErrorCode = 4130
	ErrorCodeForbidden              ErrorCode = 4030
	ErrorCodeInternalServer         ErrorCode = 5000
)

type (
	DefaultResponse[T any] struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Errors  any    `json:"errors,omitempty"`
		Data    T      `json:"data,omitempty"`
	}
)

var (
	httpStatusToErrorCodeMapper = map[int]ErrorCode{
		fiber.StatusBadRequest:            ErrorCodeBadRequest,
		fiber.StatusUnauthorized:          ErrorCodeUnauthorized,
		fiber.StatusNotFound:              ErrorCodePageNotFound,
		fiber.StatusRequestEntityTooLarge: ErrorCodeRequestEntityTooLarge,
		fiber.StatusForbidden:             ErrorCodeForbidden,
		fiber.StatusMethodNotAllowed:      ErrorCodeStatusMethodNotAllowed,
		fiber.StatusInternalServerError:   ErrorCodeInternalServer,
	}
	successMessageResponseMapper = map[int]string{
		fiber.StatusOK:      "Success",
		fiber.StatusCreated: "Created",
	}
	errorMessageResponseMapper = map[int]string{
		fiber.StatusBadRequest:            "Bad Request",
		fiber.StatusUnauthorized:          "Unauthorized",
		fiber.StatusNotFound:              "Page Not Found",
		fiber.StatusRequestEntityTooLarge: "Request Entity Too Large",
	}
)

type Error struct {
	Message    string `json:"message"`
	Code       int    `json:"code"`
	Errors     any    `json:"errors"`
	HttpStatus int    `json:"-"`
}

type SubError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

func NewSuccessResponse[T any](data T, prefix string, code int) *DefaultResponse[T] {
	return &DefaultResponse[T]{
		Code:    fmt.Sprintf("%s-%d", prefix, code),
		Message: successMessageResponseMapper[code],
		Data:    data,
	}
}

func NewError(code int) *Error {
	return &Error{
		Code:    code,
		Message: errorMessageResponseMapper[code],
		Errors:  []SubError{},
	}
}

func GetErrorCode(httpStatus int) ErrorCode {
	return httpStatusToErrorCodeMapper[httpStatus]
}

func NewValidateError(subErr interface{}) error {
	err := NewError(fiber.StatusBadRequest)
	err.HttpStatus = fiber.StatusBadRequest
	err.Errors = subErr
	return errors.WithStack(err)
}
