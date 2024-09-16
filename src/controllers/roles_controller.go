package controllers

import (
	"go-api/paginates"
	"go-api/responses"
	"go-api/src/requests"
	"go-api/src/services"
	"strconv"

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

	// Parse pagination parameters from request
	limitStr := ctx.Query("limit", "10") // Default to 10 items per page
	pageStr := ctx.Query("page", "1")    // Default to page 1

	//filter request

	filters := requests.FilterRequest{
		Search:    ctx.Query("search", ""),
		StartDate: ctx.Query("startDate", ""),
		EndDate:   ctx.Query("endDate", ""),
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid item parameter"})
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page parameter"})
	}

	paginate := paginates.PaginateRequest{
		Limit: limit,
		Page:  page,
	}

	// Call repository method
	paginatedResponse, err := c.serviceRoles.GetRoles(ctx.Context(), paginate, filters)
	if err != nil {
		return responses.NewErrorResponses(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, paginatedResponse)
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
