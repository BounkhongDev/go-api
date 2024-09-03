package repositories

import (
	"context"

	"go-api/src/models"

	"gorm.io/gorm"
)

type RolesRepository interface {
	//Insert your function interface
	// GetAll Roles
	GetRoles(ctx context.Context) ([]models.Roles, error)
	//Create Roles
	CreateRoles(ctx context.Context, roles models.Roles) error
}

type rolesRepository struct{ db *gorm.DB }

func NewRolesRepository(db *gorm.DB) RolesRepository {
	// db.Migrator().DropTable(models.Roles{})
	db.AutoMigrate(models.Roles{})
	return &rolesRepository{db: db}
}

func (r *rolesRepository) GetRoles(ctx context.Context) ([]models.Roles, error) {
	var roles []models.Roles
	if err := r.db.Order("created_at DESC").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *rolesRepository) CreateRoles(ctx context.Context, roles models.Roles) error {
	if err := r.db.Create(&roles).Error; err != nil {
		return err
	}
	return nil
}
