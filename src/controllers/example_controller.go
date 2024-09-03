package controllers

import (
"go-api/src/services"
	"github.com/gofiber/fiber/v2")

type ExampleController interface{
	PingController(ctx *fiber.Ctx) error
}

type exampleController struct {
serviceExample services.ExampleService
}

func NewExampleController(
serviceExample services.ExampleService,
//services
) ExampleController {
	return &exampleController{
serviceExample :serviceExample,
//services
}
}
func (c *exampleController) PingController(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
			"message": "pong",
	})
}