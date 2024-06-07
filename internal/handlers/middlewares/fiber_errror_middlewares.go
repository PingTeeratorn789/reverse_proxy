package middlewares

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/PingTeeratorn789/reverse_proxy/internal/core/domain"
	"github.com/PingTeeratorn789/reverse_proxy/internal/handlers/common"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(prefix string) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		var (
			rawError      = err.Error()
			internalError = domain.NewInternalServerError(rawError)
			httpStatus    = fiber.StatusInternalServerError
			response      = common.DefaultResponse[any]{
				Code:    fmt.Sprintf("%s-%v", prefix, common.GetErrorCode(httpStatus)),
				Message: internalError.Error(),
				Errors: []common.SubError{
					{Message: rawError},
				},
			}
		)
		if errorServer, ok := err.(*fiber.Error); ok {
			httpStatus = errorServer.Code
			response = common.DefaultResponse[any]{
				Code:    fmt.Sprintf("%s-%v", prefix, common.GetErrorCode(httpStatus)),
				Message: rawError,
			}
		}

		if reflect.TypeOf(err).String() == "*errors.withStack" {
			err = errors.Unwrap(err)
		}

		switch e := err.(type) {
		case *domain.Error:
			httpStatus = e.HttpStatus
			response.Code = fmt.Sprintf("%s-%v", prefix, e.Code)
			response.Message = e.Message
			response.Errors = e.Errors
		case *common.Error:
			httpStatus = e.HttpStatus
			response.Code = fmt.Sprintf("%s-%v", prefix, common.GetErrorCode(httpStatus))
			response.Message = e.Message
			response.Errors = e.Errors
		}

		return ctx.Status(httpStatus).JSON(response)
	}
}
