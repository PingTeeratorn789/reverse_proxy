package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type ErrorCode int

const (
	ErrorCodeBadRequest              ErrorCode = 4000
	ErrorCodeUnauthorized            ErrorCode = 4010
	ErrorCodePageNotFound            ErrorCode = 4040
	ErrorCodeForbidden               ErrorCode = 4030
	ErrorCodeDataNotFound            ErrorCode = 4041
	ErrorCodeUserStatusAlreadyExists ErrorCode = 4042
	ErrorCodeUserStatusNotFound      ErrorCode = 4043
	ErrorCodeInternalServer          ErrorCode = 5000
	ErrorCodeMongo                   ErrorCode = 5001
	ErrorCodeMySQL                   ErrorCode = 5002
)

type MessageStaus string

const (
	MessageStausCreated MessageStaus = "CREATED"
)

var (
	successMessageResponseMapper = map[int]string{
		fiber.StatusOK:      "Success",
		fiber.StatusCreated: "Created",
	}
	errorMessageResponseMapper = map[ErrorCode]string{
		ErrorCodeBadRequest:              "Bad Request",
		ErrorCodeUnauthorized:            "Unauthorized",
		ErrorCodePageNotFound:            "Page Not Found",
		ErrorCodeInternalServer:          "Internal Server Error",
		ErrorCodeDataNotFound:            "Not Found",
		ErrorCodeUserStatusNotFound:      "Status Not Found",
		ErrorCodeUserStatusAlreadyExists: "Status Already Exists",
	}
)

type (
	DefaultDataResponse    struct{}
	DefaultResponse[T any] struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Errors  any    `json:"errors,omitempty"`
		Data    T      `json:"data,omitempty"`
	}
)

type SubError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type Error struct {
	Message    string    `json:"message"`
	Code       ErrorCode `json:"code"`
	HttpStatus int       `json:"-"`
	Errors     any       `json:"errors"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code ErrorCode) *Error {
	return &Error{
		Code:    code,
		Message: errorMessageResponseMapper[code],
		Errors:  []SubError{},
	}
}

func NewInternalServerError(msg string) error {
	err := NewError(ErrorCodeInternalServer)
	err.HttpStatus = fiber.StatusInternalServerError
	err.Errors = []SubError{{
		Message: msg,
	}}
	return errors.WithStack(err)
}

func NewDataNotFound() error {
	return errors.WithStack(NewError(ErrorCodeDataNotFound))
}

func NewUserStatusAlreadyExists() error {
	err := NewError(ErrorCodeUserStatusAlreadyExists)
	err.HttpStatus = fiber.StatusNotFound
	return errors.WithStack(err)
}

func NewUserStatusNotfound() error {
	err := NewError(ErrorCodeUserStatusNotFound)
	err.HttpStatus = fiber.StatusNotFound
	return errors.WithStack(err)
}

func NewMongoError(msg string) error {
	err := NewError(ErrorCodeMongo)
	err.HttpStatus = fiber.StatusInternalServerError
	err.Errors = []SubError{{
		Message: msg,
	}}
	return errors.WithStack(err)
}

func NewMySQLError(msg string) error {
	err := NewError(ErrorCodeMySQL)
	err.HttpStatus = fiber.StatusInternalServerError
	err.Errors = []SubError{{
		Message: msg,
	}}
	return errors.WithStack(err)
}
