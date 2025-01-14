package response

import "github.com/gofiber/fiber/v2"

func SuccessResponse(ctx *fiber.Ctx, statusCode int, message string, data any) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
		"links": fiber.Map{
			"self": ctx.OriginalURL(),
		},
	})
}
