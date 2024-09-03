package services

import (
	"context"
	"database/sql"
	"go-api/src/models"
	"go-api/src/repositories"
	requests "go-api/src/requests"
	response "go-api/src/responses"
	"go-api/src/utils/mapper"
)

type RolesService interface {
	// Insert your function interface
	// GetAll Roles
	GetRoles(ctx context.Context) ([]response.Role, error)

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

func (s *rolesService) GetRoles(ctx context.Context) ([]response.Role, error) {
	// get roles from repository
	roles, err := s.repositoryRoles.GetRoles(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return []response.Role{}, nil
		}
		return nil, err
	}
	// concat roles data to response.Role
	newRoles := []response.Role{}

	for _, role := range roles {
		var newRole response.Role
		mapper.StructMapper(role, &newRole)
		newRole.ID = role.ID
		newRole.CreatedAt = role.CreatedAt
		newRole.UpdatedAt = role.UpdatedAt

		newRoles = append(newRoles, newRole)

	}

	return newRoles, nil
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
