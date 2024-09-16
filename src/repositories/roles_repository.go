package repositories

import (
	"context"

	"go-api/paginates"
	"go-api/src/models"
	"go-api/src/responses"

	"gorm.io/gorm"
)

type RolesRepository interface {
	//Insert your function interface
	// GetAll Roles by Paginate
	GetRoles(ctx context.Context, paginate paginates.PaginateRequest) (*paginates.PaginatedResponse, error)

	//Create Roles
	CreateRoles(ctx context.Context, roles models.Roles) error
}

type rolesRepository struct{ db *gorm.DB }

func NewRolesRepository(db *gorm.DB) RolesRepository {
	// db.Migrator().DropTable(models.Roles{})
	db.AutoMigrate(models.Roles{})
	return &rolesRepository{db: db}
}

func (r *rolesRepository) GetRoles(ctx context.Context, paginate paginates.PaginateRequest) (*paginates.PaginatedResponse, error) {
	var roles []responses.Role
	var total int64

	// Count the total number of records
	if err := r.db.Model(&responses.Roles{}).Count(&total).Error; err != nil {
		return nil, err
	}

	// Calculate offset
	offset := (paginate.Page - 1) * paginate.Limit

	// Fetch the paginated results
	if err := r.db.Limit(paginate.Limit).Offset(offset).Find(&roles).Error; err != nil {
		return nil, err
	}

	// Create IFindAndCountAll struct
	result := paginates.IFindAndCountAll{
		Count: total,
		Rows:  roles,
	}

	// Use PaginationResult to get the paginated response
	paginatedResponse := paginates.PaginationResult(paginate.Page, paginate.Limit, result)

	return &paginatedResponse, nil
}

func (r *rolesRepository) CreateRoles(ctx context.Context, roles models.Roles) error {
	if err := r.db.Create(&roles).Error; err != nil {
		return err
	}
	return nil
}
