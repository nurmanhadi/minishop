package response

import (
	"errors"
	"minishop/pkg/infrastructure/exception"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(ctx *fiber.Ctx, statusCode int, message string, err string) error {
	return ctx.Status(statusCode).JSON(&fiber.Map{
		"status":  "error",
		"message": message,
		"errors":  err,
		"links": &fiber.Map{
			"self": ctx.OriginalURL(),
		},
	})
}
func ExceptionError(ctx *fiber.Ctx, err error) error {
	if err != nil {
		if errors.Is(err, exception.ProductNotFound) {
			return ErrorResponse(ctx, 404, "not found", err.Error())
		} else if finalErr, ok := err.(*exception.ErrorValidation); ok {
			return ErrorResponse(ctx, 400, "bad request", finalErr.Message)
		} else {
			return ErrorResponse(ctx, 500, "internal server error", err.Error())
		}
	}
	return nil
}
