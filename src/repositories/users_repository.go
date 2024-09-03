package repositories

import (
	"context"
	"go-api/src/models"
	"go-api/src/responses"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsersRepository interface {
	//Insert your function interface

	// GetAll Users
	GetUsers(ctx context.Context) ([]responses.User, error)

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

func (r *usersRepository) GetUsers(ctx context.Context) ([]responses.User, error) {
	var users []models.Users
	var usersResponses []responses.User
	if err := r.db.Model(&users).Preload("Roles").Find(&usersResponses).Error; err != nil {
		return usersResponses, err
	}
	return usersResponses, nil
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
