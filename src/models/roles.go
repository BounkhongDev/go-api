package models

type Roles struct {
	DefaultModel `gorm:"embedded"`
	RoleName     *string `gorm:"unique;not null"`
}
