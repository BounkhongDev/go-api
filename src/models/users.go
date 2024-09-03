package models

import "github.com/google/uuid"

type Users struct {
	DefaultModel `gorm:"embedded"`
	Fullname     *string `gorm:"not null"`
	RolesID      uuid.UUID
	Roles        Roles `gorm:"foreignKey:RolesID;references:ID"`
}
