package routes

import (
	"go-api/src/controllers"

	"github.com/gofiber/fiber/v2"
)

type fiberRoutes struct {
	controller      controllers.ExampleController
	usersController controllers.UsersController
	rolesController controllers.RolesController
}

func (r fiberRoutes) Install(app *fiber.App) {
	route := app.Group("api/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Get("ping", r.controller.PingController)
	//get roles
	route.Get("roles", r.rolesController.GetRoles)
	//create roles
	route.Post("roles", r.rolesController.CreateRoles)

	//get users
	route.Get("users", r.usersController.GetUsers)

	//get users by id
	route.Get("users/:id", r.usersController.GetUserByID)

	//create users
	route.Post("users", r.usersController.CreateUsers)
}

func NewFiberRoutes(
	controller controllers.ExampleController,
	//new controller
	usersController controllers.UsersController,
	rolesController controllers.RolesController,
) Routes {
	return &fiberRoutes{
		controller:      controller,
		usersController: usersController,
		rolesController: rolesController,
	}
}
