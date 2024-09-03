package repositories

import (
"gorm.io/gorm"
"go-api/src/models"
)

type ExampleRepository interface{
//Insert your function interface
}

type exampleRepository struct {db *gorm.DB}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
// db.Migrator().DropTable(models.Example{})
db.AutoMigrate(models.Example{})
	return &exampleRepository{db: db}
}