package responses

import (
	"time"

	"github.com/google/uuid"
)

type Roles struct {
	ID        uuid.UUID `json:"id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Fullname  string    `json:"fullname"`
	Roles     Roles     `json:"role"`
	RolesID   uuid.UUID `json:"roles_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
