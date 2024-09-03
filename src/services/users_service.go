package services

import (
	"context"
	"go-api/paginates"
	"go-api/src/models"
	"go-api/src/repositories"
	requests "go-api/src/requests"
	response "go-api/src/responses"
	"go-api/src/utils/mapper"

	"github.com/google/uuid"
)

type UsersService interface {
	// Insert your function interface
	// GetAll Users by Paginate
	GetUsers(ctx context.Context, paginate paginates.PaginateRequest) (*paginates.PaginatedResponse, error)

	// Get User by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (response.User, error)

	// Create Users
	CreateUsers(ctx context.Context, users requests.User) error
}

type usersService struct {
	repositoryUsers repositories.UsersRepository
}

func NewUsersService(
	repositoryUsers repositories.UsersRepository,
	// repo
) UsersService {
	return &usersService{
		repositoryUsers: repositoryUsers,
		// repo
	}
}

func (s *usersService) GetUsers(ctx context.Context, paginate paginates.PaginateRequest) (*paginates.PaginatedResponse, error) {
	// get users from repository
	users, err := s.repositoryUsers.GetUsers(ctx, paginate)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *usersService) GetUserByID(ctx context.Context, id uuid.UUID) (response.User, error) {
	user, err := s.repositoryUsers.GetUserByID(ctx, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *usersService) CreateUsers(ctx context.Context, users requests.User) error {

	// concat user data to models.Users
	var newUser models.Users
	mapper.StructMapper(users, &newUser)

	if err := s.repositoryUsers.CreateUsers(ctx, newUser); err != nil {
		return err
	}
	return nil
}
