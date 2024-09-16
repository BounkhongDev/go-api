package repositories

import (
	"context"

	"go-api/paginates"
	"go-api/src/models"
	"go-api/src/requests"
	"go-api/src/responses"

	"gorm.io/gorm"
)

type RolesRepository interface {
	//Insert your function interface
	// GetAll Roles by Paginate
	GetRoles(ctx context.Context, paginate paginates.PaginateRequest, filters requests.FilterRequest) (*paginates.PaginatedResponse, error)

	//Create Roles
	CreateRoles(ctx context.Context, roles models.Roles) error
}

type rolesRepository struct{ db *gorm.DB }

func NewRolesRepository(db *gorm.DB) RolesRepository {
	// db.Migrator().DropTable(models.Roles{})
	db.AutoMigrate(models.Roles{})
	return &rolesRepository{db: db}
}

func (r *rolesRepository) GetRoles(ctx context.Context, paginate paginates.PaginateRequest, filters requests.FilterRequest) (*paginates.PaginatedResponse, error) {
	var roles []responses.Role
	var total int64

	// Start building the query
	query := r.db.Model(&responses.Role{})

	// Apply filters (if any)
	if filters.StartDate != "" && filters.EndDate != "" {
		query = query.Where("DATE(created_at) BETWEEN ? AND ?", filters.StartDate, filters.EndDate)
	}

	if filters.Search != "" {
		query = query.Where("LOWER(role_name) LIKE LOWER(?)", "%"+filters.Search+"%")
	}

	// Count the total number of records after filters are applied
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination (limit and offset)
	offset := (paginate.Page - 1) * paginate.Limit
	if err := query.Limit(paginate.Limit).Offset(offset).Find(&roles).Error; err != nil {
		return nil, err
	}

	// Prepare the result using IFindAndCountAll struct
	result := paginates.IFindAndCountAll{
		Count: total,
		Rows:  roles,
	}

	// Use PaginationResult to create the paginated response
	paginatedResponse := paginates.PaginationResult(paginate.Page, paginate.Limit, result)

	return &paginatedResponse, nil
}

func (r *rolesRepository) CreateRoles(ctx context.Context, roles models.Roles) error {
	if err := r.db.Create(&roles).Error; err != nil {
		return err
	}
	return nil
}
