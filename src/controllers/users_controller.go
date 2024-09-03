package controllers

import (
	"go-api/responses"
	"go-api/src/services"

	"go-api/src/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UsersController interface {
	//get all users
	GetUsers(ctx *fiber.Ctx) error

	//get user by id
	GetUserByID(ctx *fiber.Ctx) error

	//Create Users
	CreateUsers(ctx *fiber.Ctx) error
}

type usersController struct {
	serviceUsers services.UsersService
}

func NewUsersController(
	serviceUsers services.UsersService,
	// services
) UsersController {
	return &usersController{
		serviceUsers: serviceUsers,
		// services
	}
}

func (c *usersController) GetUsers(ctx *fiber.Ctx) error {

	defer ctx.Context().Done()

	users, err := c.serviceUsers.GetUsers(ctx.Context())
	if err != nil {
		return responses.NewErrorResponses(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, users)
}

func (c *usersController) GetUserByID(ctx *fiber.Ctx) error {

	defer ctx.Context().Done()

	id := ctx.Params("id")
	//convert string to uuid

	// string to uuid
	uid, err := uuid.Parse(id)

	users, err := c.serviceUsers.GetUserByID(ctx.Context(), uid)
	if err != nil {
		return responses.NewErrorResponses(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, users)
}

func (c *usersController) CreateUsers(ctx *fiber.Ctx) error {

	defer ctx.Context().Done()

	var users requests.User
	if err := ctx.BodyParser(&users); err != nil {
		return responses.NewErrorResponses(ctx, err)
	}

	if err := c.serviceUsers.CreateUsers(ctx.Context(), users); err != nil {
		return responses.NewErrorResponses(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, "Success Create Users")
}
