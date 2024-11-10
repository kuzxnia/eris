package exception

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/kuzxnia/eris/control-plane/pkg/exception"
	"github.com/kuzxnia/eris/control-plane/pkg/interfaces/web/dto"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(exception.ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		exception.PanicLogging(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.GeneralResponse{
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, notFoundError := err.(exception.NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(dto.GeneralResponse{
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	_, unauthorizedError := err.(exception.UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.GeneralResponse{
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	_, alreadyExistsError := err.(exception.ResourceAlreadyExistsError)
	if alreadyExistsError {
		return ctx.Status(fiber.StatusConflict).JSON(dto.GeneralResponse{
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(dto.GeneralResponse{
		Message: "General Error",
		Data:    err.Error(),
	})
}
