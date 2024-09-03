package services

import (
	"context"
	"database/sql"
	"go-api/src/models"
	"go-api/src/repositories"
	requests "go-api/src/requests"
	response "go-api/src/responses"
	"go-api/src/utils/mapper"

	"github.com/google/uuid"
)

type UsersService interface {
	// Insert your function interface
	// GetAll Users
	GetUsers(ctx context.Context) ([]response.User, error)

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

func (s *usersService) GetUsers(ctx context.Context) ([]response.User, error) {
	// get users from repository
	users, err := s.repositoryUsers.GetUsers(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return []response.User{}, nil
		}
		return nil, err
	}

	// map users data to response.User
	newUsers := []response.User{}
	for _, user := range users {
		var newUser response.User

		// map the basic fields
		mapper.StructMapper(user, &newUser)
		newUser.ID = user.ID
		newUser.CreatedAt = user.CreatedAt
		newUser.UpdatedAt = user.UpdatedAt

		// map the single role
		newUser.Roles = response.Roles{
			ID:        user.Roles.ID,        // Access the single role ID
			RoleName:  *user.Roles.RoleName, // Access the single role RoleName
			CreatedAt: user.Roles.CreatedAt, // Access the single role CreatedAt
			UpdatedAt: user.Roles.UpdatedAt, // Access the single role UpdatedAt
		}

		newUsers = append(newUsers, newUser)
	}

	return newUsers, nil
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
