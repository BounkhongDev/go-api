package responses

import (
	"go-api/errs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	code    int
	message string
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewErrorResponses(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errs.AppError:
		code = e.Status
		message = e.Message
	case error:
		code = http.StatusUnprocessableEntity
		message = err.Error()
	}
	errorResponse := ErrorResponse{
		Success: false,
		Status:  code,
		Error:   message,
	}
	return ctx.Status(code).JSON(errorResponse)
}

func NewSuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"status":  http.StatusOK,
		"data":    data,
	})
}

func NewSuccessMsg(ctx *fiber.Ctx, msg interface{}) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"status":  http.StatusOK,
		"msg":     msg,
	})
}

func NewCreateSuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"status":  http.StatusCreated,
		"data":    data,
	})
}

func NewSuccessMessage(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"status":  http.StatusOK,
		"message": data,
	})
}

func NewErrorValidate(ctx *fiber.Ctx, data interface{}) error {
	validateError := fiber.Map{
		"success": false,
		"error":   data,
		"status":  http.StatusUnprocessableEntity,
	}
	return ctx.Status(http.StatusUnprocessableEntity).JSON(validateError)
}

// validate not found
func NewErrorNotFound(ctx *fiber.Ctx, data interface{}) error {
	validateError := fiber.Map{
		"success": false,
		"error":   data,
		"status":  http.StatusNotFound,
	}
	return ctx.Status(http.StatusNotFound).JSON(validateError)
}
