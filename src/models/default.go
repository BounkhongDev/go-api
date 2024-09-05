package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DefaultModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy *uuid.UUID
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdateBy  *uuid.UUID
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy *uuid.UUID
}
