package repositories

import (
	"context"
	"go-api/src/models"
	"go-api/src/responses"

	"go-api/paginates"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsersRepository interface {
	//Insert your function interface

	// GetAll Users by Paginate
	GetUsers(ctx context.Context, paginate paginates.PaginateRequest) (*paginates.PaginatedResponse, error)

	//Get User by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (responses.User, error)

	//Create Users
	CreateUsers(ctx context.Context, users models.Users) error
}

type usersRepository struct{ db *gorm.DB }

func NewUsersRepository(db *gorm.DB) UsersRepository {
	// db.Migrator().DropTable(models.Users{})
	db.AutoMigrate(models.Users{})
	return &usersRepository{db: db}
}

func (r *usersRepository) GetUsers(ctx context.Context, paginate paginates.PaginateRequest) (*paginates.PaginatedResponse, error) {
	var users []responses.User
	var total int64

	// Count the total number of records
	if err := r.db.Model(&responses.User{}).Count(&total).Error; err != nil {
		return nil, err
	}

	// Calculate offset
	offset := (paginate.Page - 1) * paginate.Limit

	// Fetch the paginated results
	if err := r.db.Preload("Roles").
		Limit(paginate.Limit).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, err
	}

	// Create IFindAndCountAll struct
	result := paginates.IFindAndCountAll{
		Count: total,
		Rows:  users,
	}

	// Use PaginationResult to get the paginated response
	paginatedResponse := paginates.PaginationResult(paginate.Page, paginate.Limit, result)

	return &paginatedResponse, nil
}

func (r *usersRepository) GetUserByID(ctx context.Context, id uuid.UUID) (responses.User, error) {
	var users models.Users
	var usersResponses responses.User

	if err := r.db.Model(&users).Where("id", id).Preload("Roles").Find(&usersResponses).Error; err != nil {
		return usersResponses, err
	}

	return usersResponses, nil
}

func (r *usersRepository) CreateUsers(ctx context.Context, users models.Users) error {
	if err := r.db.Create(&users).Error; err != nil {
		return err
	}
	return nil
}
