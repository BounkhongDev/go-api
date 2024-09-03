package controllers

import (
	"go-api/responses"
	"go-api/src/requests"
	"go-api/src/services"

	"github.com/gofiber/fiber/v2"
)

type RolesController interface {
	//get all roles
	GetRoles(ctx *fiber.Ctx) error

	//Create Roles
	CreateRoles(ctx *fiber.Ctx) error
}

type rolesController struct {
	serviceRoles services.RolesService
}

func NewRolesController(
	serviceRoles services.RolesService,
	// services
) RolesController {
	return &rolesController{
		serviceRoles: serviceRoles,
		// services
	}
}

func (c *rolesController) GetRoles(ctx *fiber.Ctx) error {

	defer ctx.Context().Done()

	roles, err := c.serviceRoles.GetRoles(ctx.Context())
	if err != nil {
		return responses.NewErrorResponses(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, roles)
}

func (c *rolesController) CreateRoles(ctx *fiber.Ctx) error {

	defer ctx.Context().Done()

	var roles requests.Role
	if err := ctx.BodyParser(&roles); err != nil {
		return responses.NewErrorResponses(ctx, err)
	}

	if err := c.serviceRoles.CreateRoles(ctx.Context(), roles); err != nil {
		return responses.NewErrorResponses(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, "Success Create Roles")
}
