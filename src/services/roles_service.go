package services

import (
	"context"
	"go-api/paginates"
	"go-api/src/models"
	"go-api/src/repositories"
	requests "go-api/src/requests"
	"go-api/src/utils/mapper"
)

type RolesService interface {
	// Insert your function interface
	// GetAll Roles by Paginate
	GetRoles(ctx context.Context, paginate paginates.PaginateRequest) (*paginates.PaginatedResponse, error)

	// Create Roles
	CreateRoles(ctx context.Context, roles requests.Role) error
}

type rolesService struct {
	repositoryRoles repositories.RolesRepository
}

func NewRolesService(
	repositoryRoles repositories.RolesRepository,
	// repo
) RolesService {
	return &rolesService{
		repositoryRoles: repositoryRoles,
		// repo
	}
}

func (s *rolesService) GetRoles(ctx context.Context, paginate paginates.PaginateRequest) (*paginates.PaginatedResponse, error) {
	// get roles from repository
	roles, err := s.repositoryRoles.GetRoles(ctx, paginate)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *rolesService) CreateRoles(ctx context.Context, roles requests.Role) error {
	// concat role data to models.Roles
	var newRole models.Roles

	mapper.StructMapper(roles, &newRole)

	// insert user to repository
	if err := s.repositoryRoles.CreateRoles(ctx, newRole); err != nil {
		return err
	}

	return nil
}
