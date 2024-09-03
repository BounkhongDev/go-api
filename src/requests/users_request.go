package requests

import "github.com/google/uuid"

type User struct {
	Fullname string    `json:"fullname" validate:"required"`
	RolesID  uuid.UUID `json:"roles_id" validate:"required"`
}
