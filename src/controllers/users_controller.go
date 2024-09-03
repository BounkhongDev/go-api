package controllers

import (
	"go-api/paginates"
	"go-api/responses"
	"go-api/src/services"
	"strconv"

	"go-api/src/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UsersController interface {
	//get all users with pagination
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

func (uc *usersController) GetUsers(ctx *fiber.Ctx) error {
	// Parse pagination parameters from request

	limitStr := ctx.Query("limit", "10") // Default to 10 items per page
	pageStr := ctx.Query("page", "1")    // Default to page 1

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
	paginatedResponse, err := uc.serviceUsers.GetUsers(ctx.Context(), paginate)
	if err != nil {
		return responses.NewErrorResponses(ctx, err)
	}

	// return ctx.Status(fiber.StatusOK).JSON(paginatedResponse)

	return responses.NewSuccessResponse(ctx, paginatedResponse)
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
	//if user not found return 404
	if users.ID == uuid.Nil {
		return responses.NewErrorNotFound(ctx, "User not found")
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
